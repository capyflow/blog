import client from './client'

export async function loginByPwd(email:string, password:string){
  return client.post('/v1/blog/console/login/pwd', { email, password })
}

export async function publishArticle(payload: {title:string;content:string;category:string;background:string}){
  return client.post('/v1/blog/console/article/publish', payload)
}

export async function updateArticle(id:string, payload: {id:string;title:string;content:string}){
  return client.post(`/v1/blog/console/article/update/${id}`, payload)
}

export async function deleteArticle(id:string){
  return client.delete(`/v1/blog/console/article/delete/${id}`)
}
