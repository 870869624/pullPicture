<template>
  <div class="page-container">
    <div class="config-grid">
      <div class="config-card">
        <div class="card-header">
          <span class="card-icon">🌐</span>
          <h3>基础设置</h3>
        </div>
        <div class="card-body">
          <div class="config-item">
            <label class="config-label">系统名称</label>
            <input type="text" v-model="basicConfig.systemName" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">系统版本</label>
            <input type="text" v-model="basicConfig.version" class="config-input" disabled />
          </div>
          <div class="config-item">
            <label class="config-label">版权信息</label>
            <input type="text" v-model="basicConfig.copyright" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">系统描述</label>
            <textarea v-model="basicConfig.description" class="config-textarea" rows="3"></textarea>
          </div>
        </div>
      </div>
      
      <div class="config-card">
        <div class="card-header">
          <span class="card-icon">🔒</span>
          <h3>安全设置</h3>
        </div>
        <div class="card-body">
          <div class="config-item">
            <label class="config-label">登录超时时间</label>
            <div class="input-with-unit">
              <input type="number" v-model="securityConfig.timeout" class="config-input" />
              <span class="unit">分钟</span>
            </div>
          </div>
          <div class="config-item">
            <label class="config-label">密码有效期</label>
            <div class="input-with-unit">
              <input type="number" v-model="securityConfig.passwordExpire" class="config-input" />
              <span class="unit">天</span>
            </div>
          </div>
          <div class="config-item toggle-item">
            <label class="config-label">启用IP白名单</label>
            <button 
              :class="['toggle-btn', { active: securityConfig.enableWhitelist }]"
              @click="securityConfig.enableWhitelist = !securityConfig.enableWhitelist"
            >
              <span class="toggle-slider"></span>
            </button>
          </div>
          <div class="config-item toggle-item">
            <label class="config-label">启用日志记录</label>
            <button 
              :class="['toggle-btn', { active: securityConfig.enableLogging }]"
              @click="securityConfig.enableLogging = !securityConfig.enableLogging"
            >
              <span class="toggle-slider"></span>
            </button>
          </div>
        </div>
      </div>
      
      <div class="config-card">
        <div class="card-header">
          <span class="card-icon">📧</span>
          <h3>邮件设置</h3>
        </div>
        <div class="card-body">
          <div class="config-item">
            <label class="config-label">SMTP服务器</label>
            <input type="text" v-model="emailConfig.smtpServer" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">SMTP端口</label>
            <input type="number" v-model="emailConfig.smtpPort" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">发件人邮箱</label>
            <input type="email" v-model="emailConfig.senderEmail" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">发件人名称</label>
            <input type="text" v-model="emailConfig.senderName" class="config-input" />
          </div>
          <div class="config-item">
            <label class="config-label">邮箱密码</label>
            <div class="password-input">
              <input 
                :type="showEmailPassword ? 'text' : 'password'" 
                v-model="emailConfig.password" 
                class="config-input" 
              />
              <button class="toggle-password" @click="showEmailPassword = !showEmailPassword">
                {{ showEmailPassword ? '🙈' : '👁️' }}
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <div class="config-card">
        <div class="card-header">
          <span class="card-icon">🎨</span>
          <h3>外观设置</h3>
        </div>
        <div class="card-body">
          <div class="config-item">
            <label class="config-label">主题颜色</label>
            <div class="color-picker">
              <input type="color" v-model="appearanceConfig.themeColor" class="color-input" />
              <span class="color-preview" :style="{ background: appearanceConfig.themeColor }"></span>
            </div>
          </div>
          <div class="config-item">
            <label class="config-label">布局模式</label>
            <div class="radio-group">
              <label :class="['radio-option', { active: appearanceConfig.layout === 'light' }]">
                <input type="radio" v-model="appearanceConfig.layout" value="light" />
                <span>亮色模式</span>
              </label>
              <label :class="['radio-option', { active: appearanceConfig.layout === 'dark' }]">
                <input type="radio" v-model="appearanceConfig.layout" value="dark" />
                <span>暗色模式</span>
              </label>
              <label :class="['radio-option', { active: appearanceConfig.layout === 'auto' }]">
                <input type="radio" v-model="appearanceConfig.layout" value="auto" />
                <span>跟随系统</span>
              </label>
            </div>
          </div>
          <div class="config-item">
            <label class="config-label">侧边栏折叠</label>
            <button 
              :class="['toggle-btn', { active: appearanceConfig.collapseSidebar }]"
              @click="appearanceConfig.collapseSidebar = !appearanceConfig.collapseSidebar"
            >
              <span class="toggle-slider"></span>
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <div class="action-bar">
      <button class="save-btn" @click="handleSave">
        <span>💾</span>
        <span>保存设置</span>
      </button>
      <button class="reset-btn" @click="handleReset">
        <span>🔄</span>
        <span>重置默认</span>
      </button>
      <button class="export-btn" @click="handleExport">
        <span>📤</span>
        <span>导出配置</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const showEmailPassword = ref(false)

const basicConfig = reactive({
  systemName: '内部拉入系统',
  version: 'v1.0.0',
  copyright: '© 2024 内部系统. All rights reserved.',
  description: '内部管理系统，用于IP白名单管理、库存搜索、衣品上传等功能'
})

const securityConfig = reactive({
  timeout: 30,
  passwordExpire: 90,
  enableWhitelist: true,
  enableLogging: true
})

const emailConfig = reactive({
  smtpServer: 'smtp.example.com',
  smtpPort: 587,
  senderEmail: 'admin@example.com',
  senderName: '系统管理员',
  password: '******'
})

const appearanceConfig = reactive({
  themeColor: '#667eea',
  layout: 'light',
  collapseSidebar: false
})

const handleSave = () => {
  alert('设置已保存！')
}

const handleReset = () => {
  if (confirm('确定要重置为默认设置吗？')) {
    basicConfig.systemName = '内部拉入系统'
    basicConfig.version = 'v1.0.0'
    basicConfig.copyright = '© 2024 内部系统. All rights reserved.'
    basicConfig.description = '内部管理系统'
    
    securityConfig.timeout = 30
    securityConfig.passwordExpire = 90
    securityConfig.enableWhitelist = true
    securityConfig.enableLogging = true
    
    emailConfig.smtpServer = 'smtp.example.com'
    emailConfig.smtpPort = 587
    emailConfig.senderEmail = 'admin@example.com'
    emailConfig.senderName = '系统管理员'
    emailConfig.password = '******'
    
    appearanceConfig.themeColor = '#667eea'
    appearanceConfig.layout = 'light'
    appearanceConfig.collapseSidebar = false
    
    alert('已重置为默认设置')
  }
}

const handleExport = () => {
  alert('配置已导出')
}
</script>

<style scoped>
.page-container {
  width: 100%;
}

.config-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
  margin-bottom: 24px;
}

.config-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  overflow: hidden;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-bottom: 1px solid #eee;
}

.card-icon {
  font-size: 1.4rem;
}

.card-header h3 {
  font-size: 1.1rem;
  font-weight: 600;
  color: #1a1a2e;
}

.card-body {
  padding: 24px;
}

.config-item {
  margin-bottom: 20px;
}

.config-item:last-child {
  margin-bottom: 0;
}

.config-label {
  display: block;
  font-size: 0.9rem;
  font-weight: 500;
  color: #333;
  margin-bottom: 8px;
}

.config-input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.config-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.config-input:disabled {
  background: #f8f9fa;
  color: #999;
}

.config-textarea {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 0.9rem;
  resize: vertical;
  transition: all 0.3s ease;
}

.config-textarea:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-with-unit {
  display: flex;
  align-items: center;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.input-with-unit:focus-within {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-with-unit .config-input {
  border: none;
  border-radius: 0;
  padding: 12px 16px;
}

.unit {
  padding: 0 16px;
  color: #666;
  font-size: 0.9rem;
  background: #f8f9fa;
}

.toggle-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.toggle-btn {
  width: 50px;
  height: 28px;
  background: #e4e8ec;
  border: none;
  border-radius: 14px;
  cursor: pointer;
  position: relative;
  transition: all 0.3s ease;
}

.toggle-btn.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.toggle-slider {
  position: absolute;
  top: 3px;
  left: 3px;
  width: 22px;
  height: 22px;
  background: #fff;
  border-radius: 50%;
  transition: all 0.3s ease;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
}

.toggle-btn.active .toggle-slider {
  left: 25px;
}

.password-input {
  display: flex;
  align-items: center;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.password-input:focus-within {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.password-input .config-input {
  border: none;
  border-radius: 0;
  flex: 1;
}

.toggle-password {
  padding: 12px 16px;
  background: #f8f9fa;
  border: none;
  cursor: pointer;
  font-size: 1rem;
}

.color-picker {
  display: flex;
  align-items: center;
  gap: 12px;
}

.color-input {
  width: 50px;
  height: 44px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  cursor: pointer;
}

.color-preview {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  border: 2px solid #e4e8ec;
}

.radio-group {
  display: flex;
  gap: 16px;
}

.radio-option {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: #f8f9fa;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.radio-option.active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
  border: 2px solid #667eea;
}

.radio-option input {
  cursor: pointer;
}

.action-bar {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  padding: 20px;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
}

.save-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 28px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.save-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.reset-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 28px;
  background: #fff;
  border: 2px solid #e4e8ec;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.reset-btn:hover {
  background: #f5f7fa;
  border-color: #667eea;
}

.export-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 28px;
  background: #fff;
  border: 2px solid #e4e8ec;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.export-btn:hover {
  background: #f5f7fa;
  border-color: #667eea;
}
</style>