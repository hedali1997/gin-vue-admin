<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="生日:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="学校:" prop="school">
          <el-input v-model="formData.school" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学历:" prop="education">
          <el-select v-model="formData.education" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in educationOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="专业:" prop="major">
          <el-input v-model="formData.major" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="签名:" prop="reamrk">
          <el-input v-model="formData.reamrk" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CommunityUser'
}
</script>

<script setup>
import {
  createCommunityUser,
  updateCommunityUser,
  findCommunityUser
} from '@/api/communityUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const educationOptions = ref([])
const formData = ref({
  nickname: '',
  userName: '',
  password: '',
  birthday: new Date(),
  school: '',
  education: undefined,
  major: '',
  avatar: '',
  phone: '',
  email: '',
  reamrk: '',
})
// 验证规则
const rule = reactive({
  userName: [{
    required: true,
    message: '请填写用户名',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findCommunityUser({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.recommunityUser
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  educationOptions.value = await getDictFunc('education')
}

init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCommunityUser(formData.value)
        break
      case 'update':
        res = await updateCommunityUser(formData.value)
        break
      default:
        res = await createCommunityUser(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style></style>
