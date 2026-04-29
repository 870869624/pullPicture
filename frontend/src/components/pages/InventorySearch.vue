<template>
  <div class="page-container">
    <div class="page-card">
      <div class="card-header">
        <div class="header-left">
          <span class="card-icon">📦</span>
          <h2>库存搜索</h2>
        </div>
      </div>
      
      <div class="filter-section">
        <div class="filter-group">
          <label>商品名称</label>
          <input type="text" v-model="filters.name" placeholder="输入商品名称" />
        </div>
        <div class="filter-group">
          <label>商品编码</label>
          <input type="text" v-model="filters.code" placeholder="输入商品编码" />
        </div>
        <div class="filter-group">
          <label>分类</label>
          <select v-model="filters.category">
            <option value="">全部分类</option>
            <option value="shirt">衬衫</option>
            <option value="pants">裤子</option>
            <option value="dress">连衣裙</option>
            <option value="coat">外套</option>
          </select>
        </div>
        <div class="filter-group">
          <label>库存状态</label>
          <select v-model="filters.status">
            <option value="">全部状态</option>
            <option value="instock">有库存</option>
            <option value="low">库存不足</option>
            <option value="outstock">缺货</option>
          </select>
        </div>
        <button class="filter-btn" @click="handleSearch">搜索</button>
        <button class="reset-btn" @click="resetFilters">重置</button>
      </div>
      
      <div class="stats-row">
        <div class="stat-card">
          <div class="stat-icon">📊</div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.total }}</span>
            <span class="stat-label">总商品数</span>
          </div>
        </div>
        <div class="stat-card">
          <div class="stat-icon">✅</div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.instock }}</span>
            <span class="stat-label">有库存</span>
          </div>
        </div>
        <div class="stat-card">
          <div class="stat-icon">⚠️</div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.low }}</span>
            <span class="stat-label">库存不足</span>
          </div>
        </div>
        <div class="stat-card">
          <div class="stat-icon">❌</div>
          <div class="stat-info">
            <span class="stat-value">{{ stats.outstock }}</span>
            <span class="stat-label">缺货</span>
          </div>
        </div>
      </div>
      
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>商品编码</th>
              <th>商品名称</th>
              <th>分类</th>
              <th>规格</th>
              <th>库存数量</th>
              <th>状态</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in mockData" :key="item.id">
              <td>{{ item.code }}</td>
              <td>
                <div class="product-info">
                  <img :src="item.image" :alt="item.name" class="product-img" />
                  <span>{{ item.name }}</span>
                </div>
              </td>
              <td>{{ item.category }}</td>
              <td>{{ item.spec }}</td>
              <td :class="['stock-count', item.stockStatus]">{{ item.stock }}</td>
              <td>
                <span :class="['status-badge', item.stockStatus]">
                  {{ item.stockStatus === 'instock' ? '有库存' : item.stockStatus === 'low' ? '库存不足' : '缺货' }}
                </span>
              </td>
              <td>
                <button class="action-btn view" @click="handleView(item)">查看详情</button>
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
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const filters = reactive({
  name: '',
  code: '',
  category: '',
  status: ''
})

const currentPage = ref(1)
const totalPages = ref(10)

const stats = ref({
  total: 1256,
  instock: 892,
  low: 234,
  outstock: 130
})

const mockData = [
  { id: 1, code: 'SP001', name: '纯棉修身衬衫', category: '衬衫', spec: 'M/L/XL', stock: 156, stockStatus: 'instock', image: 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=white%20cotton%20shirt%20product%20photo%20minimal%20background&image_size=square' },
  { id: 2, code: 'SP002', name: '高腰阔腿裤', category: '裤子', spec: 'S/M/L', stock: 23, stockStatus: 'low', image: 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=black%20wide%20leg%20pants%20product%20photo%20minimal%20background&image_size=square' },
  { id: 3, code: 'SP003', name: '碎花连衣裙', category: '连衣裙', spec: 'M/L', stock: 0, stockStatus: 'outstock', image: 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=floral%20dress%20product%20photo%20minimal%20background&image_size=square' },
  { id: 4, code: 'SP004', name: '羊毛大衣', category: '外套', spec: 'M/L/XL/XXL', stock: 89, stockStatus: 'instock', image: 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=wool%20coat%20product%20photo%20minimal%20background&image_size=square' },
  { id: 5, code: 'SP005', name: '牛仔夹克', category: '外套', spec: 'S/M/L/XL', stock: 15, stockStatus: 'low', image: 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=denim%20jacket%20product%20photo%20minimal%20background&image_size=square' }
]

const handleSearch = () => {
  alert('搜索功能')
}

const resetFilters = () => {
  filters.name = ''
  filters.code = ''
  filters.category = ''
  filters.status = ''
}

const handleView = (item) => {
  alert(`查看商品: ${item.name}`)
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

.filter-section {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  margin-bottom: 24px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 12px;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.filter-group label {
  font-size: 0.85rem;
  font-weight: 500;
  color: #666;
}

.filter-group input,
.filter-group select {
  padding: 10px 14px;
  border: 2px solid #e4e8ec;
  border-radius: 8px;
  font-size: 0.9rem;
  min-width: 180px;
  transition: all 0.3s ease;
}

.filter-group input:focus,
.filter-group select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.filter-btn {
  align-self: flex-end;
  padding: 10px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.filter-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.reset-btn {
  align-self: flex-end;
  padding: 10px 24px;
  background: #fff;
  border: 2px solid #e4e8ec;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.reset-btn:hover {
  background: #f5f7fa;
}

.stats-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  margin-bottom: 24px;
}

.stat-card {
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid #e4e8ec;
}

.stat-icon {
  font-size: 2rem;
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 12px;
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1a1a2e;
}

.stat-label {
  font-size: 0.85rem;
  color: #666;
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

.product-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.product-img {
  width: 40px;
  height: 40px;
  object-fit: cover;
  border-radius: 8px;
}

.stock-count {
  font-weight: 600;
}

.stock-count.instock {
  color: #16a34a;
}

.stock-count.low {
  color: #d97706;
}

.stock-count.outstock {
  color: #dc2626;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.status-badge.instock {
  background: #dcfce7;
  color: #16a34a;
}

.status-badge.low {
  background: #fef3c7;
  color: #d97706;
}

.status-badge.outstock {
  background: #fee2e2;
  color: #dc2626;
}

.action-btn {
  padding: 6px 14px;
  border: none;
  border-radius: 6px;
  font-size: 0.8rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.action-btn.view {
  background: #dbeafe;
  color: #2563eb;
}

.action-btn.view:hover {
  background: #bfdbfe;
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
</style>