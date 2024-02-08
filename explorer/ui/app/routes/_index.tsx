import type { MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import { Footer } from '~/components/shared/footer'
import Navbar from '~/components/shared/navbar'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export default function Index() {
  return (
    <div>
      <Navbar />
      <div className="flex flex-col">
        <div className="grid grid-cols-2 gap-4 place-items-center m3">
          <div className="grow flex-1">
            <XBlockDataTable />
          </div>
          <div className="grow flex-1">
            <XMsgDataTable />
          </div>
        </div>
      </div>
      <Footer />
    </div>
  )
}
