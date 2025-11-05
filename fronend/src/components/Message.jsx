import { useState, useRef, useEffect } from 'react'
import MarkdownRenderer from './MarkdownRenderer'
import './Message.css'

function Message({ message, enableTypewriter = true, onAddToStory }) {
  const isUser = message.role === 'user'
  const [showMenu, setShowMenu] = useState(false)
  const menuRef = useRef(null)

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (menuRef.current && !menuRef.current.contains(event.target)) {
        setShowMenu(false)
      }
    }

    if (showMenu) {
      document.addEventListener('mousedown', handleClickOutside)
    }

    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [showMenu])

  const handleAddToStory = () => {
    // 优先使用documentId，如果没有则使用id（向后兼容）
    const documentId = message.documentId || message.id
    if (onAddToStory && documentId) {
      onAddToStory(documentId)
    }
    setShowMenu(false)
  }

  return (
    <div className={`message ${isUser ? 'message-user' : 'message-assistant'}`}>
      <div className="message-content">
        <div className="message-text">
          {isUser ? (
            <div className="user-text">{message.content}</div>
          ) : (
            <MarkdownRenderer text={message.content} speed={20} enableTypewriter={enableTypewriter} />
          )}
        </div>
        {!isUser && (
          <div className="message-actions" ref={menuRef}>
            <button
              className="message-action-btn"
              onClick={(e) => {
                e.stopPropagation()
                setShowMenu(!showMenu)
              }}
              title="更多操作"
            >
              <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                <circle cx="8" cy="8" r="1.5" fill="currentColor" />
                <circle cx="8" cy="3" r="1.5" fill="currentColor" />
                <circle cx="8" cy="13" r="1.5" fill="currentColor" />
              </svg>
            </button>
            {showMenu && (
              <div className="message-menu">
                <button
                  className="message-menu-item"
                  onClick={(e) => {
                    e.stopPropagation()
                    handleAddToStory()
                  }}
                >
                  添加为故事
                </button>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}

export default Message

