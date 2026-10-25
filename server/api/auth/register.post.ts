import { useDB } from '~~/server/utils/db'
import { hashPassword, createSession } from '~~/server/utils/auth'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const body = await readBody(event)
  
  const { name, last_name, email, password, confirm_password } = body
  
  if (!name || !last_name || !email || !password) {
    throw createError({
      statusCode: 400,
      statusMessage: 'All fields are required'
    })
  }
  
  if (password !== confirm_password) {
    throw createError({
      statusCode: 400,
      statusMessage: 'Passwords do not match'
    })
  }
  
  // Check if user exists
  const existing = await db.prepare('SELECT id FROM users WHERE email = ?').bind(email).first()
  if (existing) {
    throw createError({
      statusCode: 400,
      statusMessage: 'Email already registered'
    })
  }
  
  const hashedPassword = await hashPassword(password)
  const now = new Date().toISOString()
  
  const result = await db.prepare(
    'INSERT INTO users (name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)'
  ).bind(name, last_name, email, hashedPassword, now, now).run()
  
  const userId = result.meta.last_row_id
  
  await createSession(event, userId)
  
  const user = await db.prepare('SELECT id, name, last_name, email FROM users WHERE id = ?')
    .bind(userId).first()
  
  return { user }
})
