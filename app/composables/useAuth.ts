export const useAuth = () => {
  const user = useState('auth.user', () => null)
  
  const fetchUser = async () => {
    try {
      const data: any = await $fetch('/api/auth/me')
      user.value = data.user
      return data.user
    } catch (e) {
      user.value = null
      return null
    }
  }
  
  const login = async (email: string, password: string) => {
    const data: any = await $fetch('/api/auth/login', {
      method: 'POST',
      body: { email, password }
    })
    user.value = data.user
    return data.user
  }
  
  const register = async (name: string, last_name: string, email: string, password: string, confirm_password: string) => {
    const data: any = await $fetch('/api/auth/register', {
      method: 'POST',
      body: { name, last_name, email, password, confirm_password }
    })
    user.value = data.user
    return data.user
  }
  
  const logout = async () => {
    await $fetch('/api/auth/logout', { method: 'POST' })
    user.value = null
    await navigateTo('/login')
  }
  
  return {
    user: readonly(user),
    fetchUser,
    login,
    register,
    logout
  }
}
