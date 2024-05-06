import { Footer } from '~/components/shared/footer'
import Navbar from '~/components/shared/navbar'
import { GetXBlock } from '../../components/queries/block'
import { Link, useNavigate, useParams, useSearchParams } from '@remix-run/react'
import { BackBtn } from '~/components/details/BackBtn'
import Tag from '~/components/shared/tag'
import { From } from '~/components/details/From'
import { To } from '~/components/details/To'
import { TabList } from '~/components/details/TabList'
import { GetXMsg } from '~/components/queries/messages'

export default function Index() {
  const params = useParams()
  const [searchParams] = useSearchParams()

  // const data = GetXMsg("0x40b1", "0xc8", "0x3c5c")

//   console.log(data);

    const navigate = useNavigate();

    const onBackClickHandler = () => {
      navigate("/");
    }


  return (
    <div className="px-8 md:px-20">
      <div className="flex h-full w-full flex-col">
        {/* <Overview /> */}
        <div className={'h-5'}></div>

        <div className="w-full">
          <BackBtn onBackClickHandler={onBackClickHandler}/>

          <h3 className="mt-5">
            XMsg <span className="p-2">/</span> <span className="text-default">34-3567-345</span>
          </h3>

          <div className="mt-5 p-4 w-full bg-raised rounded-lg">
            {/* Offset */}
            <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
              <p className="w-24 sm:w-48">Offset</p>
              <p className="text-default">345</p>
            </div>
            {/* Status */}
            <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
              <p className="w-24 sm:w-48">Status</p>
              <div className='flex flex-col sm:flex-row'>
              <Tag status="FAILURE" />
              <p className="ml-5">"Reason for status"</p>
              </div>
            </div>
            {/* Data */}
            <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
              <p className="w-24 sm:w-48">Data</p>
            </div>
            <From />
            <To />
            <TabList />
          </div>
        </div>
      </div>
    </div>
  )
}
