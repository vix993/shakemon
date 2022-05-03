package e2e

// import (
// 	"ambify-wallet/internal/config"
// 	mgrpc "ambify-wallet/internal/grpc/pb"
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/rs/zerolog/log"
// 	"google.golang.org/grpc"
// )

// func Run() {
// 	<-time.After(5 * time.Second)

// 	setup()
// 	TestProtoCreateWallet(&testing.T{})
// }

// func setup() {
// 	cfg := config.GetProvider()
// 	grpcURL = fmt.Sprintf("127.0.0.1:%s", cfg.GetString(config.GRPCServerPort))
// }

// func TestProtoCreateWallet(t *testing.T) {
// 	conn, err := grpc.Dial(grpcURL, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("error connecting to grpc")
// 	}

// 	client := mgrpc.NewWalletServiceClient(conn)

// 	res, err := client.PostWallet(context.Background(), &mgrpc.PostWalletRequest{
// 		AccountId: "1",
// 		ClientId:  "1",
// 	})

// 	if err != nil {
// 		log.Fatal().Err(err).Msg("error while creating new wallet")
// 	}

// 	assert.Equal(t, res.AccountId, "1")
// 	assert.Equal(t, res.Address, "1")

// 	log.Info().Str("accountId", res.AccountId).Str("address", res.Address).Msg("wallet created successfully")
// }
