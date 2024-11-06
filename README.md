# NextBlock grpc proto spec

Use protoc to generate the grpc code from the proto file.

```bash
protoc \   
--go_out=./generated \
--go_opt=paths=source_relative \
--go-grpc_out=./generated \
--go-grpc_opt=paths=source_relative \
--proto_path ./proto/ proto/api.proto
````

## Setup

Connection prep to the grpc server with keepalive and credentials.

```go
var kacp = keepalive.ClientParameters{
Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
PermitWithoutStream: true,             // send pings even without active streams
}

type ApiCredentials struct {
    Token string
}

func (p *ApiCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
    return map[string]string{
        "authorization": fmt.Sprintf("%s", p.Token),
    }, nil
}

func (p *ApiCredentials) RequireTransportSecurity() bool {
    return false
}


func grpcConnect(address string, plaintext bool) *grpc.ClientConn {
	var opts []grpc.DialOption
	if plaintext {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		pool, _ := x509.SystemCertPool()
		creds := credentials.NewClientTLSFromCert(pool, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	opts = append(opts, grpc.WithKeepaliveParams(kacp))
	opts = append(opts, grpc.WithPerRPCCredentials(&ApiCredentials{Token: "<api key>"}))
	log.Println("Starting grpc client, connecting to", address)
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}
```

## Usage

Init the grpc connection and the api client.
Send a sample transaction.

```go
var (   
    cc = grpcConnect("fra.nextblock.io", false)
	apiClient = api.NewApiClient(cc)
)

func demo() {
	f := false
	cli := rpc.New("<rpc address>")
	bh, _ := cli.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	pk := solana.MustPrivateKeyFromBase58("<private key>")
	txBuilder := solana.NewTransactionBuilder()
	txBuilder.SetFeePayer(pk.PublicKey())
	txBuilder.SetRecentBlockHash(bh.Value.Blockhash)
	txBuilder.AddInstruction(
		system.NewTransferInstruction(1000000, pk.PublicKey(), solana.MustPublicKeyFromBase58("nEXTBLockYgngeRmRrjDV31mGSekVPqZoMGhQEZtPVG")).Build(),
	)
	tx, err := txBuilder.Build()
	if err != nil {
		log.Println(err)
		return
	}
	tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		return &pk
	})
	tb64, _ := tx.ToBase64()
	response, err := apiClient.PostSubmitV2(context.TODO(), &api.PostSubmitRequest{
		Transaction:            &api.TransactionMessage{Content: tb64},
		SkipPreFlight:          true,
		UseStakedRPCs:          &f,
		FrontRunningProtection: &f,
		FastBestEffort:         &f,
	})
	if err != nil {
		return
	}
	log.Println(response)
}
```

## Randomize wallet selection

We recommend to ranomize the wallet selection to avoid spamming the same wallet and potentially
hitting per slot wallet limits.

```go
var NextblockTipWallets = []string{
	"NEXTbLoCkB51HpLBLojQfpyVAMorm3zzKg7w9NFdqid",
	"nextBLoCkPMgmG8ZgJtABeScP35qLa2AMCNKntAP7Xc",
	"NextbLoCkVtMGcV47JzewQdvBpLqT9TxQFozQkN98pE",
	"NexTbLoCkWykbLuB1NkjXgFWkX9oAtcoagQegygXXA2",
	"NeXTBLoCKs9F1y5PJS9CKrFNNLU1keHW71rfh7KgA1X",
	"NexTBLockJYZ7QD7p2byrUa6df8ndV2WSd8GkbWqfbb",
	"neXtBLock1LeC67jYd1QdAa32kbVeubsfPNTJC1V5At",
	"nEXTBLockYgngeRmRrjDV31mGSekVPqZoMGhQEZtPVG",
}

func RandomWallet() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano())),
    return NextblockTipWallets[rand.Intn(len(NextblockTipWallets))]
}
```

## Priority fees

The minimum priority fee is `0.001` SOL or `1000000` lamports.
Higher tip will result in higher priority and better landing times.