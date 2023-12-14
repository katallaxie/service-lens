import { Table } from './components/table'
import { AddSolution } from './components/add-solution'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { DataTableProvider } from '@/components/data-table-context'

export type PageProps = {
  children?: React.ReactNode
}

export default function Page({ children }: PageProps) {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Solutions
          <SubNavSubtitle>Design, discuss, review, and build.</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddSolution />
        </SubNavActions>
      </SubNav>
      <section>
        <DataTableProvider>
          <Table />
        </DataTableProvider>
      </section>
    </>
  )
}
