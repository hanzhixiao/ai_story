# Grandma 前端

## 简介
Grandma 前端使用 React 和 Vite 开发，提供类似GPT的用户界面和打字机效果。

## 功能特性
- GPT风格的聊天界面
- 打字机效果展示AI响应
- 模型选择功能
- 流式响应接收
- 响应式设计

## 环境要求
- Node.js 16 或更高版本
- npm 或 yarn

## 安装和运行

1. 安装依赖
```bash
cd fronend
npm install
```

2. 启动开发服务器
```bash
npm run dev
```

前端将在 `http://localhost:3000` 启动

## 构建

```bash
npm run build
```

构建产物将在 `dist` 目录中

## 项目结构
```
src/
  ├── components/      # React组件
  │   ├── ChatContainer.jsx    # 聊天容器
  │   ├── MessageList.jsx       # 消息列表
  │   ├── Message.jsx           # 单条消息
  │   ├── Typewriter.jsx        # 打字机效果组件
  │   ├── InputArea.jsx         # 输入区域
  │   └── Sidebar.jsx           # 侧边栏
  ├── App.jsx          # 主应用组件
  ├── main.jsx        # 入口文件
  └── index.css       # 全局样式
```

## 使用说明
1. 在侧边栏选择要使用的大模型
2. 在输入框中输入你的想法或问题
3. 按 Enter 或点击发送按钮
4. AI的响应将以打字机效果逐字显示

