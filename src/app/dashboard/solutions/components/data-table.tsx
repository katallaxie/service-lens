'use client'

import { useDataTableContext } from './data-context'
import { columns } from './data-columns'
import { DataTable } from '@/components/data-table'

export default function LensesDataTable() {
  const dataTableContext = useDataTableContext()

  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable
        rows={dataTableContext.state.rows}
        columns={columns}
        onPaginationChange={dataTableContext.onPaginationChange}
        isFetching={dataTableContext.state.cursor.fetching}
        pagination={dataTableContext.state.cursor}
      />
    </div>
  )
}
