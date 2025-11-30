import axios from 'axios'
import { useAuth } from '../store/auth'

const client = axios.create({ baseURL: '/' })

client.interceptors.request.use(config => {
  const auth = useAuth()
  if(auth.token){
    config.headers = config.headers || {}
    config.headers.Authorization = `Beader ${auth.token}`
  }
  return config
})

export default client
