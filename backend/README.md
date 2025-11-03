# Grandma 后端

## 简介
Grandma 后端使用 Golang 和 Gin 框架开发，提供大模型API接入和流式响应功能。

## 功能特性
- 支持多种大模型API（OpenAI、Anthropic）
- 流式响应支持
- RESTful API接口
- CORS跨域支持

## 环境要求
- Go 1.21 或更高版本

## 安装和运行

1. 安装依赖
```bash
cd backend
go mod download
```

2. 配置环境变量
创建 `.env` 文件（参考 `.env.example`），配置API密钥：
```env
OPENAI_API_KEY=your_openai_api_key_here
ANTHROPIC_API_KEY=your_anthropic_api_key_here
PORT=8080
```

3. 运行服务器
```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动

## API接口

### POST /api/chat
发送聊天请求，获取流式响应

**请求体：**
```json
{
  "model": "openai",
  "message": "你好"
}
```

**响应：**
流式文本响应 (text/event-stream)

### GET /api/models
获取可用模型列表

**响应：**
```json
{
  "models": [
    {
      "id": "openai",
      "name": "GPT-3.5 Turbo",
      "provider": "OpenAI"
    },
    {
      "id": "anthropic",
      "name": "Claude 3.5 Sonnet",
      "provider": "Anthropic"
    }
  ]
}
```

### GET /health
健康检查接口

## 支持的模型
- OpenAI: GPT-3.5 Turbo
- Anthropic: Claude 3.5 Sonnet

