import { useDB } from '~~/server/utils/db'

export default defineEventHandler(async (event) => {
  const db = useDB(event)
  const user = event.context.user
  const body = await readBody(event)
  
  const { weight, entered_date } = body
  
  if (!weight || !entered_date) {
    throw createError({
      statusCode: 400,
      statusMessage: 'All fields are required'
    })
  }
  
  const now = new Date().toISOString()
  
  const result = await db.prepare(
    'INSERT INTO weights (user_id, weight, entered_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?)'
  ).bind(user.id, weight, entered_date, now, now).run()
  
  const entry = await db.prepare('SELECT * FROM weights WHERE id = ?')
    .bind(result.meta.last_row_id).first()
  
  return { weight: entry }
})
