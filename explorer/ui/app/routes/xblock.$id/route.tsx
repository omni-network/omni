import { Footer } from '~/components/shared/footer'
import Navbar from '~/components/shared/navbar'
import { GetXBlock } from '../../components/queries/block'
import { useParams, useSearchParams } from '@remix-run/react'

export default function Index() {
  const params = useParams()
  const [searchParams] = useSearchParams()

  const data = GetXBlock(params.id || '', searchParams.get('height') || '')
  return (
    <div>
      <Navbar />
      <div className="flex">
        <div className="grow"></div>
        <div className="flex-auto w-full max-w-screen-xl grid grid-cols-2 gap-4 place-items-stretch m3">
          {/* <div className={'card w-96 bg-base-100 shadow-xl'}> */}
          <div className={'card-body'}>
            <span className="bg-custom-color">XBlock Hash:</span> {data?.xblock?.BlockHash || ''}
          </div>
          {/* </div> */}
        </div>
        <div className="grow"></div>
      </div>
      <Footer />
    </div>
  )
}
