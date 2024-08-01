import { useChat } from "@/services/mutations";
import { Message } from "@/services/types";
import { Bot, MessageSquare, X } from "lucide-react";
import { useState } from "react";

const AssistantUI = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [messages, setMessages] = useState<Message[]>([]);
  const [message, setMessage] = useState("");

  const handleToggle = () => setIsOpen(!isOpen);

  const chatMutation = useChat();

  const handleSend = () => {
    const newMessage = { role: "user", content: message };
    const updatedMessages = [...messages, newMessage];
    setMessages(updatedMessages);

    chatMutation.mutate(updatedMessages, {
      onSuccess: (data) => {
        const botMessage = data.message;
        setMessages((prevMessages) => [...prevMessages, botMessage]);
      },
      onError: (error) => {
        console.error("Error:", error);
      },
    });

    setMessage("");
  };

  return (
    <div className="container">
      <div className="fixed bottom-28 right-10">
        {isOpen && (
          <div className="mt-4 p-2 bg-white rounded-lg shadow-lg w-[480px]">
            <div className="bg-gradient-to-r from-rose-900 to-rose-700 p-3 rounded-t-lg text-white flex items-center">
              <div className="flex justify-between w-full mr-4">
                <div className="flex items-center">
                  <Bot size={32} />
                  <div className="ml-6">
                    <h4 className="text-base">Assistant</h4>
                    <p className="text-sm">How can I help you?</p>
                  </div>
                </div>
                <button
                  className="bg-rose-800 hover:bg-rose-600 p-2 shadow-lg hover:shadow-sm text-white rounded-lg px-4"
                  onClick={() => setMessages([])}
                >
                  Clear
                </button>
              </div>
            </div>

            <div className="flex flex-col-reverse overflow-y-auto h-[420px] mt-2 p-2">
              {messages.map((msg, index) => (
                <div
                  key={index}
                  className={`p-2 my-1 rounded-lg ${msg.role === "user"
                    ? "bg-gray-200 self-end"
                    : "bg-rose-500 text-white  self-start"
                    }`}
                >
                  {msg.content}
                </div>
              ))}
            </div>

            <div className="flex items-center bg-gradient-to-r from-rose-900 to-rose-700 p-2 rounded-b-lg">
              <input
                type="text"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={(e) => e.key === "Enter" && handleSend()}
                className="flex-1 p-2 rounded-lg"
                placeholder="Ask a query?"
              />
              <button
                onClick={handleSend}
                className="ml-2 px-4 bg-rose-800 hover:bg-rose-600 p-2 rounded-lg text-white"
              >
                Send
              </button>
            </div>
          </div>
        )}
      </div>
      <div className="fixed bottom-10 right-10">
        <button
          onClick={handleToggle}
          className="p-4 bg-rose-800 hover:bg-rose-700 text-white rounded-full shadow-lg"
        >
          {isOpen ? <X size={32} /> : <MessageSquare size={32} />}
        </button>
      </div>
    </div>
  );
};

export default AssistantUI;
