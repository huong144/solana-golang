package Kafka

import "context"

func Main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	produce(ctx)
	//consume(ctx)
}
