<template>
  <div class="app">
    <template v-if="!isLoggedIn">
      <div class="login-container">
        <div class="background"></div>
        <div class="content">
          <h1 class="title">内部拉入系统</h1>
          <div class="login-card">
            <div class="card-header">
              <h2>欢迎登录</h2>
              <p>请输入您的账号信息</p>
            </div>
            <form @submit.prevent="handleLogin" class="login-form">
              <div class="form-group">
                <label for="username">用户名</label>
                <div class="input-wrapper">
                  <span class="icon">👤</span>
                  <input 
                    type="text" 
                    id="username" 
                    v-model="username" 
                    placeholder="请输入用户名"
                    required
                  />
                </div>
              </div>
              <div class="form-group">
                <label for="password">密码</label>
                <div class="input-wrapper">
                  <span class="icon">🔒</span>
                  <input 
                    type="password" 
                    id="password" 
                    v-model="password" 
                    placeholder="请输入密码"
                    required
                  />
                </div>
              </div>
              <button type="submit" class="login-btn" :disabled="loading">
                <span v-if="loading" class="spinner"></span>
                {{ loading ? '登录中...' : '登 录' }}
              </button>
            </form>
            <div v-if="error" class="error-message">{{ error }}</div>
            <div v-if="success" class="success-message">{{ success }}</div>
          </div>
        </div>
      </div>
    </template>
    
    <template v-else>
      <div class="main-layout">
        <Sidebar :activeMenu="activeMenu" @select="handleMenuSelect" @logout="handleLogout" />
        <ContentArea :activeMenu="activeMenu" :username="currentUser" />
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from './components/Sidebar.vue'
import ContentArea from './components/ContentArea.vue'

const isLoggedIn = ref(false)
const username = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const success = ref('')
const activeMenu = ref('ip-whitelist')
const currentUser = ref('')

const handleLogin = async () => {
  error.value = ''
  success.value = ''
  loading.value = true

  try {
    const response = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value
      })
    })

    const data = await response.json()

    if (response.ok) {
      success.value = `登录成功！欢迎, ${data.user.username}`
      currentUser.value = data.user.username
      setTimeout(() => {
        isLoggedIn.value = true
        success.value = ''
      }, 1000)
    } else {
      error.value = data.error || '登录失败'
    }
  } catch (err) {
    error.value = '网络错误，请稍后重试'
    console.error('登录错误:', err)
  } finally {
    loading.value = false
  }
}

const handleMenuSelect = (menuId) => {
  activeMenu.value = menuId
}

const handleLogout = () => {
  isLoggedIn.value = false
  username.value = ''
  password.value = ''
  activeMenu.value = 'ip-whitelist'
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app {
  min-height: 100vh;
}

.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  animation: gradientShift 15s ease infinite;
  background-size: 200% 200%;
}

@keyframes gradientShift {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.background::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 20% 80%, rgba(255, 255, 255, 0.15) 0%, transparent 50%),
              radial-gradient(circle at 80% 20%, rgba(255, 255, 255, 0.1) 0%, transparent 50%),
              radial-gradient(circle at 40% 40%, rgba(255, 255, 255, 0.08) 0%, transparent 40%);
}

.content {
  position: relative;
  z-index: 1;
  text-align: center;
  padding: 20px;
}

.title {
  font-family: 'Ma Shan Zheng', cursive;
  font-size: 4rem;
  color: #ffffff;
  text-shadow: 2px 2px 10px rgba(0, 0, 0, 0.2);
  margin-bottom: 40px;
  letter-spacing: 10px;
  animation: float 3s ease-in-out infinite;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

.login-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 420px;
  animation: slideUp 0.6s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card-header h2 {
  font-size: 1.8rem;
  color: #1a1a2e;
  margin-bottom: 8px;
}

.card-header p {
  color: #666;
  font-size: 0.9rem;
  margin-bottom: 30px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  text-align: left;
}

.form-group label {
  display: block;
  color: #333;
  font-weight: 500;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

.input-wrapper {
  display: flex;
  align-items: center;
  background: #f5f5f5;
  border-radius: 12px;
  padding: 12px 16px;
  transition: all 0.3s ease;
  border: 2px solid transparent;
}

.input-wrapper:focus-within {
  background: #fff;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-wrapper .icon {
  font-size: 1.2rem;
  margin-right: 12px;
  color: #999;
}

.input-wrapper input {
  flex: 1;
  border: none;
  outline: none;
  background: transparent;
  font-size: 1rem;
  color: #333;
}

.input-wrapper input::placeholder {
  color: #aaa;
}

.login-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  padding: 14px 24px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.login-btn:active:not(:disabled) {
  transform: translateY(0);
}

.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.error-message {
  margin-top: 16px;
  padding: 12px;
  background: #fee2e2;
  color: #dc2626;
  border-radius: 8px;
  font-size: 0.9rem;
  animation: fadeIn 0.3s ease;
}

.success-message {
  margin-top: 16px;
  padding: 12px;
  background: #dcfce7;
  color: #16a34a;
  border-radius: 8px;
  font-size: 0.9rem;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.main-layout {
  display: flex;
  min-height: 100vh;
}

@media (max-width: 480px) {
  .title {
    font-size: 2.5rem;
    letter-spacing: 5px;
    margin-bottom: 30px;
  }
  
  .login-card {
    padding: 30px 20px;
  }
}
</style>