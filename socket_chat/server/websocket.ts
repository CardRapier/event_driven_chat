import { Server } from "socket.io";
import http from "http";

// Create a basic HTTP server
const server = http.createServer();

// Attach Socket.IO to the HTTP server
const io = new Server(server, {
	cors: {
		origin: "*", // In production, restrict this to specific domains
	},
});

// Handle new socket connections
io.on("connection", (socket) => {
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

// Start the server
const PORT = 3000;
server.listen(PORT, () => {
	console.log(`Server is listening on port ${PORT}`);
});
