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
      <div className="flex">
        <div className="grow"></div>
        <div className="flex-auto w-full max-w-screen-xl grid grid-cols-2 gap-4 place-items-stretch m3">
          <div className="flex-auto">
            <XBlockDataTable />
          </div>
          <div className="flex-auto">
            <XMsgDataTable />
          </div>
        </div>
        <div className="grow"></div>
      </div>
      <Footer />
    </div>
  )
}
