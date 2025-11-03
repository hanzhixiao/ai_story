import Typewriter from './Typewriter'
import './Message.css'

function Message({ message }) {
  const isUser = message.role === 'user'

  return (
    <div className={`message ${isUser ? 'message-user' : 'message-assistant'}`}>
      <div className="message-content">
        {isUser ? (
          <div className="message-text">{message.content}</div>
        ) : (
          <div className="message-text">
            <Typewriter text={message.content} speed={20} />
          </div>
        )}
      </div>
    </div>
  )
}

export default Message

