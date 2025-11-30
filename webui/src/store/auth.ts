import { defineStore } from 'pinia'

export const useAuth = defineStore('auth', {
  state: () => ({ token: (typeof localStorage !== 'undefined' ? localStorage.getItem('token') || '' : '') as string }),
  actions: {
    setToken(t: string){
      this.token = t
      if (typeof localStorage !== 'undefined') localStorage.setItem('token', t)
    },
    clearToken(){
      this.token = ''
      if (typeof localStorage !== 'undefined') localStorage.removeItem('token')
    }
  }
})
