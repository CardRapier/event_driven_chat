import { ChatMessage, ChatSentiment, SentimentAnalysis } from "../types.ts";

import WebsocketIO from "../server/websocket.ts";
import kafkajs from "kafkajs";

export default class Consumer {
	socket: WebsocketIO;
	consumer: kafkajs.Consumer;
	kafka: kafkajs.Kafka;
	constructor(websocket: WebsocketIO) {
		this.socket = websocket;
		// Initialize the Kafka client for Redpanda (Kafka-compatible)
		this.kafka = new kafkajs.Kafka({
			clientId: "websocket_consumer-1",
			brokers: ["100.105.26.112:9092"], // Replace with your Redpanda broker(s)
		});

		// Create a consumer and assign it to a consumer group
		this.consumer = this.kafka.consumer({
			groupId: "websocket_consumers",
		});
	}

	public async initConnectionConsumer() {
		// Connect to the broker
		await this.consumer.connect();
		console.log("Connected to Redpanda!");

		// Subscribe to the topic. Set `fromBeginning` to true if needed.
		await this.consumer.subscribe({ topic: "chat", fromBeginning: false });
	}

	public async consume() {
		// Run the consumer and process each message
		await this.consumer.run({
			eachMessage: async ({
				topic,
				partition,
				message,
			}: kafkajs.EachMessagePayload) => {
				const key: string | null = message.key ? message.key.toString() : null;
				let value: string | null = message.value
					? message.value.toString()
					: null;

				console.log(`Received message: topic=${topic}, partition=${partition}`);
				console.log(`Key: ${key}`);
				console.log(`Value: ${value}`);

				if (!value) {
					return;
				}

				value = JSON.parse(value);

				if (key === "chat_message") {
					const chatMessage = ChatMessage.fromJSON(value);
					this.socket.emitNewMessage(chatMessage);
				} else if (key === "sentiment_message") {
					const chatSentiment = ChatSentiment.fromJSON(value);
					this.socket.emitMessageSentiment(chatSentiment);
				}
			},
		});
	}
}
