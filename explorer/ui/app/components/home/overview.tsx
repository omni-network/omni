import React from 'react'

interface Props {}

const Card: React.FC<{
  data: { title: string; value: number }
}> = ({ data }) => {
  const { title, value } = data
  return (
    <div className="h-[140px] bg-raised rounded-lg px-6 py-3">
      <span className="text-subtlest mb-[5px] block"> {title}</span>
      <h4 className="text-default">{value}</h4>
    </div>
  )
}

const Overview: React.FC<Props> = ({}) => {
  const cards = [
    { title: 'Total XMsgs', value: 79489200 },
    { title: 'Xblock Count', value: 15489200 },
    { title: 'Total Receipts', value: 4289200 },
    { title: 'Total Pending', value: 539200 },
  ]
  return (
    <div className="mb-12 mt-8">
      <h5 className="text-default">Omni X-Explorer</h5>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3 mt-3">
        {cards.map((card, i) => (
          <Card key={`card-${i}`} data={card} />
        ))}
      </div>
    </div>
  )
}

export default Overview
