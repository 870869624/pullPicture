<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="logo">
        <span class="logo-icon">🚀</span>
        <span class="logo-text">内部系统</span>
      </div>
    </div>
    <nav class="sidebar-nav">
      <ul>
        <li 
          v-for="item in menuItems" 
          :key="item.id"
          :class="{ active: activeMenu === item.id }"
          @click="selectMenu(item.id)"
        >
          <span class="menu-icon">{{ item.icon }}</span>
          <span class="menu-text">{{ item.name }}</span>
          <span v-if="activeMenu === item.id" class="active-indicator"></span>
        </li>
      </ul>
    </nav>
    <div class="sidebar-footer">
      <button class="logout-btn" @click="handleLogout">
        <span>👤</span>
        <span>退出登录</span>
      </button>
    </div>
  </aside>
</template>

<script setup>
defineProps({
  activeMenu: {
    type: String,
    default: 'ip-whitelist'
  }
})

const emit = defineEmits(['select', 'logout'])

const menuItems = [
  { id: 'ip-whitelist', name: 'IP白名单', icon: '🛡️' },
  { id: 'inventory-search', name: '库存搜索', icon: '📦' },
  { id: 'clothing-upload', name: '衣品上传', icon: '👔' },
  { id: 'system-config', name: '系统配置', icon: '⚙️' }
]

const selectMenu = (id) => {
  emit('select', id)
}

const handleLogout = () => {
  emit('logout')
}
</script>

<style scoped>
.sidebar {
  width: 260px;
  background: linear-gradient(180deg, #1a1a2e 0%, #16213e 100%);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  box-shadow: 4px 0 20px rgba(0, 0, 0, 0.15);
}

.sidebar-header {
  padding: 30px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-icon {
  font-size: 1.8rem;
  animation: bounce 2s ease-in-out infinite;
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}

.logo-text {
  font-size: 1.3rem;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.sidebar-nav {
  flex: 1;
  padding: 20px 0;
}

.sidebar-nav ul {
  list-style: none;
  padding: 0 10px;
}

.sidebar-nav li {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  margin-bottom: 8px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  color: #a0aec0;
}

.sidebar-nav li:hover {
  background: rgba(102, 126, 234, 0.15);
  color: #fff;
  transform: translateX(4px);
}

.sidebar-nav li.active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.3) 0%, rgba(118, 75, 162, 0.3) 100%);
  color: #fff;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.2);
}

.menu-icon {
  font-size: 1.3rem;
}

.menu-text {
  font-size: 0.95rem;
  font-weight: 500;
}

.active-indicator {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 4px;
  height: 24px;
  background: linear-gradient(180deg, #667eea 0%, #764ba2 100%);
  border-radius: 2px 0 0 2px;
}

.sidebar-footer {
  padding: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.logout-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 12px 16px;
  background: rgba(239, 68, 68, 0.15);
  border: none;
  border-radius: 10px;
  color: #ef4444;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  background: rgba(239, 68, 68, 0.25);
  transform: translateY(-2px);
}
</style>