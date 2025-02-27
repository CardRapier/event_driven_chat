import Consumer from "./events/consumer.ts";
import WebsocketIO from "./server/websocket.ts";

const webSocket = new WebsocketIO();
const consumer = new Consumer(webSocket);
await consumer.initConnectionConsumer();
await consumer.consume();
