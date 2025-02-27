import kafkajs from "kafkajs";

async function createConsumer(): Promise<void> {
	// Initialize the Kafka client for Redpanda (Kafka-compatible)
	const kafka = new kafkajs.Kafka({
		clientId: "websocket_consumer-1",
		brokers: ["100.105.26.112:9092"], // Replace with your Redpanda broker(s)
	});

	// Create a consumer and assign it to a consumer group
	const consumer: kafkajs.Consumer = kafka.consumer({ groupId: "websocket_consumers" });

	// Connect to the broker
	await consumer.connect();
	console.log("Connected to Redpanda!");

	// Subscribe to the topic. Set `fromBeginning` to true if needed.
	await consumer.subscribe({ topic: "chat", fromBeginning: false });

	// Run the consumer and process each message
	await consumer.run({
		eachMessage: async ({ topic, partition, message }: kafkajs.EachMessagePayload) => {
			const key: string | null = message.key ? message.key.toString() : null;
			const value: string | null = message.value
				? message.value.toString()
				: null;

			console.log(`Received message: topic=${topic}, partition=${partition}`);
			console.log(`Key: ${key}`);
			console.log(`Value: ${value}`);
		},
	});
}

// Start the consumer and catch any errors.
createConsumer().catch((error) => {
	console.error("Error in consumer:", error);
});
