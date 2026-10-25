import { useDB } from '~~/server/utils/db'
import { verifyPassword, createSession } from '~~/server/utils/auth'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const body = await readBody(event)
  
  const { email, password } = body
  
  if (!email || !password) {
    throw createError({
      statusCode: 400,
      statusMessage: 'Email and password are required'
    })
  }
  
  const user: any = await db.prepare('SELECT * FROM users WHERE email = ?').bind(email).first()
  
  if (!user) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Invalid credentials'
    })
  }
  
  const valid = await verifyPassword(password, user.password)
  if (!valid) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Invalid credentials'
    })
  }
  
  await createSession(event, user.id)
  
  return { 
    user: {
      id: user.id,
      name: user.name,
      last_name: user.last_name,
      email: user.email
    }
  }
})
