<template>
  <div class="page-container">
    <div class="page-card">
      <div class="card-header">
        <div class="header-left">
          <span class="card-icon">👔</span>
          <h2>衣品上传</h2>
        </div>
      </div>
      
      <div class="upload-area" @click="triggerUpload" @dragover.prevent @drop.prevent="handleDrop">
        <div class="upload-icon">📤</div>
        <p class="upload-title">点击或拖拽文件到此处上传</p>
        <p class="upload-desc">支持 JPG、PNG、GIF 格式，单文件不超过 10MB</p>
        <input 
          type="file" 
          ref="fileInput" 
          class="file-input" 
          accept="image/*" 
          multiple
          @change="handleFileSelect"
        />
      </div>
      
      <div v-if="uploadingFiles.length > 0" class="upload-list">
        <h3 class="list-title">上传列表</h3>
        <div v-for="(file, index) in uploadingFiles" :key="index" class="upload-item">
          <div class="file-info">
            <span class="file-icon">🖼️</span>
            <div class="file-details">
              <span class="file-name">{{ file.name }}</span>
              <span class="file-size">{{ formatSize(file.size) }}</span>
            </div>
          </div>
          <div class="upload-progress">
            <div class="progress-bar" :style="{ width: file.progress + '%' }"></div>
          </div>
          <span :class="['upload-status', file.status]">
            {{ file.status === 'uploading' ? '上传中...' : file.status === 'done' ? '✓ 完成' : '✗ 失败' }}
          </span>
        </div>
      </div>
      
      <div class="form-section">
        <h3 class="section-title">商品信息</h3>
        <div class="form-grid">
          <div class="form-group">
            <label>商品名称 *</label>
            <input type="text" v-model="formData.name" placeholder="请输入商品名称" required />
          </div>
          <div class="form-group">
            <label>商品编码 *</label>
            <input type="text" v-model="formData.code" placeholder="请输入商品编码" required />
          </div>
          <div class="form-group">
            <label>分类 *</label>
            <select v-model="formData.category" required>
              <option value="">请选择分类</option>
              <option value="shirt">衬衫</option>
              <option value="pants">裤子</option>
              <option value="dress">连衣裙</option>
              <option value="coat">外套</option>
              <option value="skirt">半身裙</option>
              <option value="accessory">配饰</option>
            </select>
          </div>
          <div class="form-group">
            <label>品牌</label>
            <input type="text" v-model="formData.brand" placeholder="请输入品牌名称" />
          </div>
          <div class="form-group">
            <label>规格</label>
            <input type="text" v-model="formData.spec" placeholder="如: S/M/L/XL" />
          </div>
          <div class="form-group">
            <label>颜色</label>
            <input type="text" v-model="formData.color" placeholder="如: 黑色/白色/红色" />
          </div>
          <div class="form-group full-width">
            <label>商品描述</label>
            <textarea v-model="formData.description" placeholder="请输入商品描述" rows="4"></textarea>
          </div>
          <div class="form-group">
            <label>价格 *</label>
            <div class="price-input">
              <span class="price-symbol">¥</span>
              <input type="number" v-model="formData.price" placeholder="0.00" required />
            </div>
          </div>
          <div class="form-group">
            <label>库存数量 *</label>
            <input type="number" v-model="formData.stock" placeholder="0" required />
          </div>
        </div>
      </div>
      
      <div class="submit-section">
        <button class="submit-btn" @click="handleSubmit">
          <span>📥</span>
          <span>提交上传</span>
        </button>
        <button class="draft-btn" @click="handleSaveDraft">保存草稿</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const fileInput = ref(null)
const uploadingFiles = ref([])

const formData = reactive({
  name: '',
  code: '',
  category: '',
  brand: '',
  spec: '',
  color: '',
  description: '',
  price: '',
  stock: ''
})

const triggerUpload = () => {
  fileInput.value?.click()
}

const handleFileSelect = (event) => {
  const files = Array.from(event.target.files)
  files.forEach(file => {
    uploadingFiles.value.push({
      name: file.name,
      size: file.size,
      progress: 0,
      status: 'uploading'
    })
    
    simulateUpload(uploadingFiles.value.length - 1)
  })
}

const handleDrop = (event) => {
  const files = Array.from(event.dataTransfer.files)
  files.forEach(file => {
    if (file.type.startsWith('image/')) {
      uploadingFiles.value.push({
        name: file.name,
        size: file.size,
        progress: 0,
        status: 'uploading'
      })
      
      simulateUpload(uploadingFiles.value.length - 1)
    }
  })
}

const simulateUpload = (index) => {
  const interval = setInterval(() => {
    if (uploadingFiles.value[index].progress < 100) {
      uploadingFiles.value[index].progress += Math.random() * 20
      if (uploadingFiles.value[index].progress >= 100) {
        uploadingFiles.value[index].progress = 100
        uploadingFiles.value[index].status = 'done'
        clearInterval(interval)
      }
    }
  }, 200)
}

const formatSize = (bytes) => {
  if (bytes < 1024) return bytes + ' B'
  if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
  return (bytes / (1024 * 1024)).toFixed(1) + ' MB'
}

const handleSubmit = () => {
  if (!formData.name || !formData.code || !formData.category || !formData.price || !formData.stock) {
    alert('请填写必填项')
    return
  }
  alert('提交成功！')
}

const handleSaveDraft = () => {
  alert('草稿已保存')
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

.upload-area {
  border: 2px dashed #e4e8ec;
  border-radius: 16px;
  padding: 48px 24px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #fafbfc 0%, #f5f7fa 100%);
  margin-bottom: 24px;
}

.upload-area:hover {
  border-color: #667eea;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
}

.upload-area.dragover {
  border-color: #667eea;
  transform: scale(1.02);
}

.upload-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.upload-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 8px;
}

.upload-desc {
  font-size: 0.9rem;
  color: #666;
}

.file-input {
  display: none;
}

.upload-list {
  margin-bottom: 24px;
}

.list-title {
  font-size: 1rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
}

.upload-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 12px;
  margin-bottom: 12px;
}

.file-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.file-icon {
  font-size: 1.5rem;
}

.file-details {
  display: flex;
  flex-direction: column;
}

.file-name {
  font-size: 0.9rem;
  font-weight: 500;
  color: #333;
}

.file-size {
  font-size: 0.8rem;
  color: #999;
}

.upload-progress {
  flex: 2;
  height: 8px;
  background: #e4e8ec;
  border-radius: 4px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
  border-radius: 4px;
  transition: width 0.2s ease;
}

.upload-status {
  font-size: 0.85rem;
  font-weight: 500;
  min-width: 80px;
}

.upload-status.uploading {
  color: #667eea;
}

.upload-status.done {
  color: #16a34a;
}

.upload-status.error {
  color: #dc2626;
}

.form-section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 1rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 20px;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group.full-width {
  grid-column: 1 / -1;
}

.form-group label {
  font-size: 0.9rem;
  font-weight: 500;
  color: #333;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 12px 16px;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-group textarea {
  resize: vertical;
}

.price-input {
  display: flex;
  align-items: center;
  border: 2px solid #e4e8ec;
  border-radius: 10px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.price-input:focus-within {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.price-symbol {
  padding: 12px 8px 12px 16px;
  color: #666;
  font-weight: 600;
}

.price-input input {
  border: none;
  padding: 12px 16px;
  flex: 1;
}

.submit-section {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.submit-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.draft-btn {
  padding: 14px 32px;
  background: #fff;
  border: 2px solid #e4e8ec;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.draft-btn:hover {
  background: #f5f7fa;
  border-color: #667eea;
}
</style>