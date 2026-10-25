import type { D1Database } from '@cloudflare/workers-types'
import Database from 'better-sqlite3'
import { join } from 'path'
import { readFileSync } from 'fs'

let localDb: any = null

function createD1Compat(db: any) {
  return {
    prepare(sql: string) {
      const stmt = db.prepare(sql)
      return {
        bind(...params: any[]) {
          return {
            async first() {
              return stmt.get(...params) || null
            },
            async all() {
              return { results: stmt.all(...params) }
            },
            async run() {
              const info = stmt.run(...params)
              return { 
                success: true, 
                meta: { last_row_id: info.lastInsertRowid, changes: info.changes } 
              }
            }
          }
        },
        async first() {
          return stmt.get() || null
        },
        async all() {
          return { results: stmt.all() }
        },
        async run() {
          const info = stmt.run()
          return { 
            success: true, 
            meta: { last_row_id: info.lastInsertRowid, changes: info.changes } 
          }
        }
      }
    }
  }
}

export function useDB(event?: any): any {
  // Try Cloudflare binding first
  const env = event?.context?.cloudflare?.env || (globalThis as any).__env
  if (env?.DB) {
    return env.DB as D1Database
  }
  
  // Fallback for local dev
  if (!localDb) {
    try {
      localDb = new Database(join(process.cwd(), 'dev.db'))
      localDb.pragma('journal_mode = WAL')
      // Initialize schema
      const schema = readFileSync(join(process.cwd(), 'schema.sql'), 'utf8')
      localDb.exec(schema)
    } catch (e) {
      console.error('Failed to init local DB:', e)
      throw createError({
        statusCode: 500,
        statusMessage: 'Database not available'
      })
    }
  }
  return createD1Compat(localDb)
}
