'use client'

import { useDataTableContext } from './data-context'
import { columns } from './data-columns'
import { DataTable } from '@/components/data-table'
import type { Lens } from '@/db/models/lens'

export default function LensesDataTable() {
  const dataTableContext = useDataTableContext()

  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable<Lens>
        columns={columns}
        onPaginationChange={dataTableContext.onPaginationChange}
        state={dataTableContext.state}
      />
    </div>
  )
}
