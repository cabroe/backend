import { useState, useEffect } from 'react'
import { ShiftType } from '../types'

export const useShifts = () => {
  const [shifts, setShifts] = useState<ShiftType[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchShifts()
  }, [])

  const fetchShifts = async () => {
    // API-Logik hier
  }

  return { shifts, loading }
}
