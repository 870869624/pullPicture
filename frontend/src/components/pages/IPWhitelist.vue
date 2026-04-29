<template>
  <div class="page-container">
    <div class="page-card">
      <div class="card-header">
        <div class="header-left">
          <span class="card-icon">🛡️</span>
          <h2>IP白名单管理</h2>
        </div>
        <button class="add-btn" @click="showAddModal = true">
          <span>+</span>
          <span>添加IP</span>
        </button>
      </div>
      
      <div class="search-bar">
        <input 
          type="text" 
          v-model="searchQuery" 
          placeholder="搜索IP地址..."
          class="search-input"
        />
        <button class="search-btn">🔍</button>
      </div>
      
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>IP地址</th>
              <th>备注</th>
              <th>状态</th>
              <th>添加时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in mockData" :key="item.id">
              <td>{{ item.ip }}</td>
              <td>{{ item.note }}</td>
              <td>
                <span :class="['status-badge', item.status]">
                  {{ item.status === 'enabled' ? '启用' : '禁用' }}
                </span>
              </td>
              <td>{{ item.time }}</td>
              <td>
                <button class="action-btn edit" @click="handleEdit(item)">编辑</button>
                <button class="action-btn delete" @click="handleDelete(item)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <div class="pagination">
        <button class="page-btn" :disabled="currentPage === 1">上一页</button>
        <span class="page-info">第 {{ currentPage }} / {{ totalPages }} 页</span>
        <button class="page-btn" :disabled="currentPage === totalPages">下一页</button>
      </div>
    </div>
    
    <div v-if="showAddModal" class="modal-overlay" @click.self="showAddModal = false">
      <div class="modal-content">
        <div class="modal-header">
          <h3>添加IP白名单</h3>
          <button class="close-btn" @click="showAddModal = false">×</button>
        </div>
        <form @submit.prevent="handleAdd" class="modal-form">
          <div class="form-group">
            <label>IP地址</label>
            <input type="text" v-model="newIP" placeholder="请输入IP地址" required />
          </div>
          <div class="form-group">
            <label>备注</label>
            <input type="text" v-model="newNote" placeholder="请输入备注信息" />
          </div>
          <div class="modal-footer">
            <button type="button" class="cancel-btn" @click="showAddModal = false">取消</button>
            <button type="submit" class="confirm-btn">确定</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const searchQuery = ref('')
const currentPage = ref(1)
const totalPages = ref(5)
const showAddModal = ref(false)
const newIP = ref('')
const newNote = ref('')

const mockData = [
  { id: 1, ip: '192.168.1.100', note: '办公室A区', status: 'enabled', time: '2024-01-15 10:30' },
  { id: 2, ip: '192.168.1.101', note: '办公室B区', status: 'enabled', time: '2024-01-15 10:35' },
  { id: 3, ip: '192.168.1.102', note: '测试环境', status: 'disabled', time: '2024-01-16 09:20' },
  { id: 4, ip: '10.0.0.50', note: '服务器集群', status: 'enabled', time: '2024-01-17 14:45' },
  { id: 5, ip: '10.0.0.51', note: '备用服务器', status: 'enabled', time: '2024-01-18 11:00' }
]

const handleEdit = (item) => {
  alert(`编辑: ${item.ip}`)
}

const handleDelete = (item) => {
  if (confirm(`确定删除IP: ${item.ip}?`)) {
    alert('删除成功')
  }
}

const handleAdd = () => {
  alert(`添加IP: ${newIP.value}`)
  showAddModal.value = false
  newIP.value = ''
  newNote.value = ''
}
</script>

<style scoped>
.page-container {
  width: 100%;
}

.page-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  padding: 24px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #eee;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.card-icon {
  font-size: 1.5rem;
}

.card-header h2 {
  font-size: 1.3rem;
  font-weight: 600;
  color: #1a1a2e;
}

.add-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.add-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.search-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
}

.search-input {
  flex: 1;
  max-width: 300px;
  padding: 12px 16px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.search-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.search-btn {
  padding: 12px 20px;
  background: #f5f7fa;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.search-btn:hover {
  background: #e4e8ec;
}

.table-container {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th,
.data-table td {
  padding: 14px 16px;
  text-align: left;
  border-bottom: 1px solid #f0f0f0;
}

.data-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #666;
  font-size: 0.85rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.data-table tbody tr:hover {
  background: #f8f9fa;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.status-badge.enabled {
  background: #dcfce7;
  color: #16a34a;
}

.status-badge.disabled {
  background: #fef3c7;
  color: #d97706;
}

.action-btn {
  padding: 6px 14px;
  border: none;
  border-radius: 6px;
  font-size: 0.8rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-right: 8px;
}

.action-btn.edit {
  background: #e0e7ff;
  color: #4338ca;
}

.action-btn.edit:hover {
  background: #c7d2fe;
}

.action-btn.delete {
  background: #fee2e2;
  color: #dc2626;
}

.action-btn.delete:hover {
  background: #fecaca;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  margin-top: 24px;
}

.page-btn {
  padding: 10px 20px;
  background: #fff;
  border: 2px solid #e4e8ec;
  border-radius: 8px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.page-btn:hover:not(:disabled) {
  background: #f5f7fa;
  border-color: #667eea;
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-info {
  font-size: 0.9rem;
  color: #666;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.modal-content {
  background: #fff;
  border-radius: 16px;
  width: 100%;
  max-width: 450px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2);
  animation: slideUp 0.3s ease;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid #eee;
}

.modal-header h3 {
  font-size: 1.2rem;
  font-weight: 600;
  color: #1a1a2e;
}

.close-btn {
  width: 32px;
  height: 32px;
  border: none;
  background: #f5f7fa;
  border-radius: 8px;
  font-size: 1.5rem;
  color: #666;
  cursor: pointer;
  transition: all 0.3s ease;
}

.close-btn:hover {
  background: #e4e8ec;
}

.modal-form {
  padding: 24px;
}

.modal-form .form-group {
  margin-bottom: 20px;
}

.modal-form label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #333;
  font-size: 0.9rem;
}

.modal-form input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.modal-form input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid #eee;
  margin-top: 12px;
}

.cancel-btn {
  padding: 10px 24px;
  background: #f5f7fa;
  border: none;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cancel-btn:hover {
  background: #e4e8ec;
}

.confirm-btn {
  padding: 10px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.confirm-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}
</style>