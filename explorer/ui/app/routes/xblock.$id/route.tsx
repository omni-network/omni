import { Footer } from '~/components/shared/footer'
import Navbar from '~/components/shared/navbar'

export default function Index() {
  return (
    <div>
      <Navbar />
      <div className="flex">
        <div className="grow"></div>
        <div className="flex-auto w-full max-w-screen-xl grid grid-cols-2 gap-4 place-items-stretch m3">
          xblock data to go here...
        </div>
        <div className="grow"></div>
      </div>
      <Footer />
    </div>
  )
}