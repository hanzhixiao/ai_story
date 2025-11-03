import { useState, useRef, useEffect } from 'react'
import MessageList from './MessageList'
import InputArea from './InputArea'
import './ChatContainer.css'

function ChatContainer({ messages, onSendMessage, isLoading }) {
  const containerRef = useRef(null)

  useEffect(() => {
    // 自动滚动到底部
    if (containerRef.current) {
      containerRef.current.scrollTop = containerRef.current.scrollHeight
    }
  }, [messages])

  return (
    <div className="chat-container">
      <div className="chat-messages" ref={containerRef}>
        <MessageList messages={messages} isLoading={isLoading} />
      </div>
      <InputArea onSendMessage={onSendMessage} isLoading={isLoading} />
    </div>
  )
}

export default ChatContainer

