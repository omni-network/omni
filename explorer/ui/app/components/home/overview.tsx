import { useLoaderData } from '@remix-run/react';
import React from 'react'
import { XmsgResponse } from '~/routes/_index';
import { ArrowIconSmall } from '../svg/arrowIconSmall';

interface Props {}

const Card: React.FC<{
  data: { title: string; value: number; sourceChainLogo: string; destLogo: string }
}> = ({ data }) => {
  const { title, value, sourceChainLogo, destLogo } = data
  return (
    <div className="h-[140px] bg-raised rounded-lg px-6 py-3  flex flex-col justify-between">
      <span className="text-subtlest text-sm block h-[26px]">{title}</span>
      <h4 className="text-default text-2xl">{value}</h4>
      <div className='flex'>
        <img className='w-5 h-5 mr-2' src={sourceChainLogo}/>
        <div className='dark:hidden'>
          <ArrowIconSmall />
        </div>
        <img className='w-5 h-5 ml-2'  src={destLogo}/>
      </div>
    </div>
  )
}

const Overview: React.FC<Props> = ({}) => {

  const data = useLoaderData<XmsgResponse>().chainStats?.stats


  const cards = [
    { title: 'Total XMsgs', value: data.totalMsgs },
  ]
  const otherThreeCards = data.topStreams.map(cardData => {
    return {
      title: `Total xmsgs from ${cardData.sourceChain.name} to ${cardData.destChain.name}`,
      value: cardData.msgCount,
      sourceChainLogo: cardData.sourceChain.logoUrl,
      destLogo: cardData.destChain.logoUrl,
    }
  })

  return (
    <div className="mb-12 mt-8">
      <h5 className="text-default">Omni X-Explorer</h5>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3 mt-3">
        <div className="h-[140px] bg-raised rounded-lg px-6 py-3 flex flex-col justify-between">
        <span className="text-subtlest text-sm block h-[26px]">{'Total XMsgs'}</span>
        <h4 className="text-default text-2xl">{data.totalMsgs}</h4>
        <span className='h-5 w-5'></span>
      </div>
        {otherThreeCards.map((card, i) => (
          <Card key={`card-${i}`} data={card} />
        ))}
      </div>
    </div>
  )
}

export default Overview
