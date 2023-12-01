import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { SidebarNav } from './components/sidebar-nav'
import { Separator } from '@/components/ui/separator'
import { GeneralForm } from './components/general-form'

type PageProps = {
  children?: React.ReactNode
}

const sidebarNavItems = [
  {
    title: 'General',
    href: '/dashboard/settings'
  },
  {
    title: 'Environments',
    href: '/dashboard/settings/environments'
  },
  {
    title: 'Appearance',
    href: '/dashboard/settings/appearance'
  }
]

export default function Layout({ children }: PageProps) {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Settings
          <SubNavSubtitle>
            Manage the settings of the service lens.
          </SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions></SubNavActions>
      </SubNav>
      <main className="p-8">
        <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
          <aside className="-mx-4 lg:w-1/5">
            <SidebarNav items={sidebarNavItems} />
          </aside>
          <div className="flex-1 lg:max-w-2xl">
            <div className="space-y-6">{children}</div>
          </div>
        </div>
      </main>
    </>
  )
}
