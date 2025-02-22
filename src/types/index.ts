export interface User {
  id: number
  name: string
  email: string
  departmentId: number
}

export interface Department {
  id: number
  name: string
}

export interface ShiftType {
  id: number
  name: string
  description?: string
  duration?: number
}
