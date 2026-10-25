import { useDB } from './db'
import bcrypt from 'bcryptjs'

export function generateSessionId(): string {
  const bytes = new Uint8Array(32)
  crypto.getRandomValues(bytes)
  return Array.from(bytes).map(b => b.toString(16).padStart(2, '0')).join('')
}

export async function hashPassword(password: string): Promise<string> {
  return bcrypt.hash(password, 10)
}

export async function verifyPassword(password: string, hash: string): Promise<boolean> {
  return bcrypt.compare(password, hash)
}

export async function createSession(event: any, userId: number) {
  const db = useDB(event)
  const sessionId = generateSessionId()
  const now = Math.floor(Date.now() / 1000)
  const ipAddress = getRequestIP(event) || null
  const userAgent = getRequestHeader(event, 'user-agent') || null
  
  await db.prepare(
    'INSERT INTO sessions (id, user_id, ip_address, user_agent, payload, last_activity) VALUES (?, ?, ?, ?, ?, ?)'
  ).bind(sessionId, userId, ipAddress, userAgent, JSON.stringify({ userId }), now).run()
  
  setCookie(event, 'health_session', sessionId, {
    httpOnly: true,
    secure: process.env.NODE_ENV === 'production',
    sameSite: 'lax',
    maxAge: 60 * 60 * 24 * 7, // 7 days
    path: '/'
  })
  
  return sessionId
}

export async function getSessionFromDB(event: any) {
  const db = useDB(event)
  const sessionId = getCookie(event, 'health_session')
  if (!sessionId) return null
  
  const session = await db.prepare(
    'SELECT * FROM sessions WHERE id = ?'
  ).bind(sessionId).first()
  
  if (!session) return null
  
  // Update last activity
  const now = Math.floor(Date.now() / 1000)
  await db.prepare('UPDATE sessions SET last_activity = ? WHERE id = ?')
    .bind(now, sessionId).run()
  
  return session
}

export async function getCurrentUser(event: any) {
  const session: any = await getSessionFromDB(event)
  if (!session || !session.user_id) return null
  
  const db = useDB(event)
  const user = await db.prepare(
    'SELECT id, name, last_name, email, created_at, updated_at FROM users WHERE id = ?'
  ).bind(session.user_id).first()
  
  return user
}

export async function destroySession(event: any) {
  const db = useDB(event)
  const sessionId = getCookie(event, 'health_session')
  if (sessionId) {
    await db.prepare('DELETE FROM sessions WHERE id = ?').bind(sessionId).run()
  }
  deleteCookie(event, 'health_session', { path: '/' })
}
