package main

import (
	"solana-crawl-service/Solana/services"
	"solana-crawl-service/cmd/stringer"
)

func main() {
	//services.FetchAllNft()
	//Solana.GetOwnerNFt("5j86xUYZEbsPowxRmr493z8CdkUbPhDa4m6W8KBfYEsy")
	//Solana.Create5000()
	services.GetSyncBlock()
	// services.TestGetBlock(151620266)
	//startCLI()
	//services.GetMetaData("5N53CHtmPHfHpzfGoJQkie6GZ1NA3Lem7x2eyqJDyi1U")
	//services.ExampleFromGetTransaction("2rjuDbsVFHv3K34CsiQUjnFYVPtBQ6QTKAYgpzyaokNycg814NNn4uS238yafGPbiRS55sx3fKWuVrh9LfbVpaTi")
	//services.GetSyncBlock()
	//services.GetAccountInfo("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	//EVM.ConnectEVM()
	//EVM.TransactionCount("0x9f83b08d90eeda539f7e2797fed3f6996917bba8")
	//Solana.FetchAllNft()
	//Kafka.TopicCrated()
	//Kafka.StartKafka()
	//Kafka.Main()
	//Database.Connect()
	//schema.ConnectDB()

}

func startCLI() {
	stringer.Execute()
}
