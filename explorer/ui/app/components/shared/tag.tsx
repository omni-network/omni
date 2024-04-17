import React from 'react'

interface Props {
  status: 'Success' | 'Failure' | 'Pending'
}

const Tag: React.FC<Props> = ({ status }) => {
  const [color, setColor] = React.useState({})

  const getColor = () => {
    switch (status) {
      case 'Success':
        return {
          text: 'text-positive',
          icon: 'text-icon-positive',
          bg: 'bg-positive',
        }
      case 'Pending':
        return {
          text: 'text-secondary',
          icon: 'text-icon-secondary',
          bg: 'bg-secondary',
        }
      case 'Failure':
        return {
          text: 'text-critical',
          icon: 'text-icon-critical',
          bg: 'bg-critical',
        }
    }
  }

  return (
    <div
      className={`py-[3.5px] px-[5.5px] text-btn-xs rounded-[4px] inline-block ${getColor()?.text} ${getColor()?.bg}`}
    >
      {status}
    </div>
  )
}

export default Tag
