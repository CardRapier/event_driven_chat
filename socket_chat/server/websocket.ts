import { Server } from "socket.io";
import http from "http";

export default class WebsocketIO {
	private server: http.Server;
	private io: Server;

	constructor() {
		this.server = http.createServer();
		this.io = new Server(this.server, {
			cors: {
				origin: "*",
			},
		});

		this.initSocketEvents();

		const PORT = 3000;

		this.server.listen(PORT, () => {
			console.log(`Server is listening on port ${PORT}`);
		});
	}

	private initSocketEvents() {
		this.io.on("connection", (socket) => {
			console.log("A new client connected");

			// Listen for events from the client
			socket.on("message", (data) => {
				console.log("Received message:", data);
				// Optionally, emit a response back to the client
				socket.emit("reply", "Message received!");
			});

			// Handle client disconnections
			socket.on("disconnect", () => {
				console.log("Client disconnected");
			});
		});
	}
}
