import Chat from "./Chat";
import io from "socket.io-client";
import useChatStore from "./stores/chat.store";
// App.js
import { useEffect } from "react";

const App = () => {
	const { addMessage, setSocket } = useChatStore();

	useEffect(() => {
		// Connect to the Socket.IO server (adjust URL as necessary)
		const socket = io("http://localhost:3000");

		// Optionally, store the socket instance in Zustand
		setSocket(socket);

		// Listen for incoming messages
		socket.on("message", (data) => {
			addMessage(data);
		});

		// Clean up on component unmount
		return () => {
			socket.off("message");
			socket.disconnect();
		};
	}, [addMessage, setSocket]);

	return (
		<div>
			<h1>Real-Time Chat with Socket.IO & Zustand</h1>
			<Chat />
		</div>
	);
};

export default App;
