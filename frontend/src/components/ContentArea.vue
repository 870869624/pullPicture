<template>
  <main class="content-area">
    <header class="content-header">
      <div class="header-info">
        <h1 class="page-title">{{ currentPage.title }}</h1>
        <p class="page-desc">{{ currentPage.description }}</p>
      </div>
      <div class="header-actions">
        <div class="user-info">
          <span class="user-icon">👤</span>
          <span class="user-name">{{ username }}</span>
        </div>
        <div class="date-display">{{ currentDate }}</div>
      </div>
    </header>
    
    <div class="content-body">
      <transition name="fade" mode="out-in">
        <component :is="currentComponent" :key="activeMenu" />
      </transition>
    </div>
  </main>
</template>

<script setup>
import { ref, computed, defineAsyncComponent } from 'vue'

const props = defineProps({
  activeMenu: {
    type: String,
    default: 'ip-whitelist'
  },
  username: {
    type: String,
    default: '用户'
  }
})

const pages = {
  'ip-whitelist': { title: 'IP白名单管理', description: '管理系统访问的IP白名单列表' },
  'inventory-search': { title: '库存搜索', description: '查询和管理库存信息' },
  'clothing-upload': { title: '衣品上传', description: '上传和管理衣品数据' },
  'system-config': { title: '系统配置', description: '配置系统参数和设置' }
}

const currentDate = computed(() => {
  const date = new Date()
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const weekDays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
  const weekDay = weekDays[date.getDay()]
  return `${year}-${month}-${day} ${weekDay}`
})

const currentPage = computed(() => pages[props.activeMenu] || pages['ip-whitelist'])

const components = {
  'ip-whitelist': defineAsyncComponent(() => import('./pages/IPWhitelist.vue')),
  'inventory-search': defineAsyncComponent(() => import('./pages/InventorySearch.vue')),
  'clothing-upload': defineAsyncComponent(() => import('./pages/ClothingUpload.vue')),
  'system-config': defineAsyncComponent(() => import('./pages/SystemConfig.vue'))
}

const currentComponent = computed(() => components[props.activeMenu] || components['ip-whitelist'])
</script>

<style scoped>
.content-area {
  flex: 1;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%);
  display: flex;
  flex-direction: column;
}

.content-header {
  padding: 24px 32px;
  background: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-info .page-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1a1a2e;
  margin-bottom: 4px;
}

.header-info .page-desc {
  font-size: 0.9rem;
  color: #666;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 24px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 10px;
}

.user-icon {
  font-size: 1.2rem;
}

.user-name {
  font-size: 0.9rem;
  font-weight: 600;
  color: #333;
}

.date-display {
  font-size: 0.85rem;
  color: #999;
}

.content-body {
  flex: 1;
  padding: 24px 32px;
  overflow-y: auto;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>