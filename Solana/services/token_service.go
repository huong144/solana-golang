package services

import (
	"github.com/portto/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBytes([]byte{91, 26, 25, 23, 127, 231, 122, 22, 0, 250, 65, 114, 15, 59, 186, 165, 25, 101, 254, 54, 37, 118, 164, 45, 84, 11, 109, 79, 187, 154, 166, 179, 182, 2, 106, 80, 22, 146, 250, 58, 216, 86, 45, 244, 34, 172, 229, 181, 14, 201, 92, 32, 204, 71, 171, 78, 21, 3, 178, 115, 34, 99, 21, 28})

var alice, _ = types.AccountFromBytes([]byte{112, 65, 32, 158, 51, 23, 136, 41, 144, 87, 104, 179, 174, 208, 230, 253, 197, 210, 155, 130, 238, 127, 39, 135, 100, 50, 151, 196, 42, 142, 222, 142, 173, 211, 79, 0, 204, 26, 80, 127, 208, 50, 26, 9, 94, 141, 229, 20, 212, 7, 64, 20, 206, 195, 135, 50, 103, 150, 114, 41, 231, 29, 102, 97})

//func CreateNFT() {
//	c := client.NewClient("https://sota-solana-local.sotatek.works")
//	// create an mint account
//	mint := types.NewAccount()
//	fmt.Println("mint:", mint.PublicKey.ToBase58())
//
//	// get init balance
//	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(
//		context.Background(),
//		tokenprog.MintAccountSize,
//	)
//	if err != nil {
//		log.Fatalf("get min balacne for rent exemption, err: %v", err)
//	}
//
//	res, err := c.GetRecentBlockhash(context.Background())
//	if err != nil {
//		log.Fatalf("get recent block hash error, err: %v\n", err)
//	}
//	tx, err := types.NewTransaction(types.NewTransactionParam{
//		Message: types.NewMessage(types.NewMessageParam{
//			FeePayer:        feePayer.PublicKey,
//			RecentBlockhash: res.Blockhash,
//			Instructions: []types.Instruction{
//				sysprog.CreateAccount(sysprog.CreateAccountParam{
//					From:     feePayer.PublicKey,
//					New:      mint.PublicKey,
//					Owner:    common.TokenProgramID,
//					Lamports: rentExemptionBalance,
//					Space:    tokenprog.MintAccountSize,
//				}),
//				tokenprog.InitializeMint(tokenprog.InitializeMintParam{
//					Decimals:   0,
//					Mint:       mint.PublicKey,
//					MintAuth:   alice.PublicKey,
//					FreezeAuth: nil,
//				}),
//			},
//		}),
//		Signers: []types.Account{feePayer, mint},
//	})
//	if err != nil {
//		log.Fatalf("generate tx error, err: %v\n", err)
//	}
//
//	txhash, err := c.SendTransaction(context.Background(), tx)
//	if err != nil {
//		log.Fatalf("send tx error, err: %v\n", err)
//	}
//
//	log.Println("txhash:", txhash)
//	_ = saveNft(models.Token{
//		Address:     mint.PublicKey.String(),
//		CreatedHash: txhash,
//	})
//}
//
//func saveNft(data models.Token) *mongo.InsertOneResult {
//	tokenRepository, ctx, client := Database.ConnectDatabase("Token")
//	c, err := tokenRepository.InsertOne(ctx, data)
//	if err != nil {
//		fmt.Printf("loi roi ne", err)
//	}
//	defer client.Disconnect(ctx)
//	return c
//}
//
//func Create5000() {
//	for i := 0; i <= 5000; i++ {
//		go CreateNFT()
//		time.Sleep(200 * time.Millisecond)
//	}
//}
//
//func handleGetOwnerNFT(client *rpc.Client, address string, index int) {
//	pubkey := solana.MustPublicKeyFromBase58(address)
//	LargestAccount, error := client.GetTokenLargestAccounts(context.TODO(), pubkey, rpc.CommitmentFinalized)
//	// Owner
//	if error != nil {
//		log.Println("Error has been detected ", error)
//	} else {
//		if len(LargestAccount.Value) > 0 {
//			ownerPub := LargestAccount.Value[0].Address
//			owner, err := client.GetAccountInfo(context.TODO(), ownerPub)
//			if err != nil {
//				log.Println("Error 2 ", err)
//			} else {
//				dataParde, _ := tokenprog.TokenAccountFromData(owner.Value.Data.GetBinary())
//				Owner := dataParde.Owner.String()
//				log.Println(Owner)
//			}
//		} else {
//			log.Println("NFT ", address, " has been no owner !", index)
//		}
//	}
//}
//
//func GetRealOwnerNFT(data []models.Token) {
//	c := rpc.New("https://sota-solana-local.sotatek.works")
//
//	for i := 0; i < len(data); i++ {
//		go handleGetOwnerNFT(c, data[i].Address, i)
//		time.Sleep(10 * time.Millisecond)
//	}
//}
//
//func FetchAllNft() {
//	// get all nft of db
//	tokenRepository, ctx, client := Database.ConnectDatabase("Token")
//	nfts, error := tokenRepository.Find(ctx, bson.M{})
//	if error != nil {
//		log.Println("co loi ne ", error)
//	}
//	var episodesFiltered []models.Token
//	if error = nfts.All(ctx, &episodesFiltered); error != nil {
//		log.Fatal(error)
//	}
//	GetRealOwnerNFT(episodesFiltered)
//	defer client.Disconnect(ctx)
//
//}
