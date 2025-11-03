import Message from './Message'
import './MessageList.css'

function MessageList({ messages, isLoading }) {
  if (messages.length === 0) {
    return (
      <div className="message-list-empty">
        <div className="welcome-message">
          <h1>欢迎使用 Grandma</h1>
          <p>AI故事创作助手，让我们开始创作吧！</p>
        </div>
      </div>
    )
  }

  return (
    <div className="message-list">
      {messages.map((message) => (
        <Message key={message.id} message={message} />
      ))}
      {isLoading && messages[messages.length - 1]?.role === 'assistant' && (
        <div className="typing-indicator">
          <span></span>
          <span></span>
          <span></span>
        </div>
      )}
    </div>
  )
}

export default MessageList

