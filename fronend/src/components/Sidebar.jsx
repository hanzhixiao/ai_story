import './Sidebar.css'

function Sidebar({ models, selectedModel, onModelChange, onNewChat }) {
  return (
    <div className="sidebar">
      <div className="sidebar-header">
        <button onClick={onNewChat} className="new-chat-button">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
            <path
              d="M8 1v14M1 8h14"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
            />
          </svg>
          新对话
        </button>
      </div>

      <div className="sidebar-content">
        <div className="model-selector">
          <label className="model-label">选择模型</label>
          <select
            value={selectedModel}
            onChange={(e) => onModelChange(e.target.value)}
            className="model-select"
          >
            {models.map((model) => (
              <option key={model.id} value={model.id}>
                {model.name}
              </option>
            ))}
          </select>
        </div>

        <div className="sidebar-footer">
          <div className="app-info">
            <p className="app-name">Grandma</p>
            <p className="app-version">v0.1</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Sidebar

