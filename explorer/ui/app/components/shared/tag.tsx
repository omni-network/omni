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
          iconClass: 'icon-check-1',
          bg: 'bg-positive',
        }
      case 'Pending':
        return {
          text: 'text-moderate',
          icon: 'text-icon-moderate',
          bg: 'bg-moderate',
          iconClass: 'icon-clock',
        }
      case 'Failure':
        return {
          text: 'text-critical',
          icon: 'text-icon-critical',
          bg: 'bg-critical',
          iconClass: 'icon-error---filled',
        }
      default:
        return {
          text: 'text-positive',
          icon: 'text-icon-positive',
          iconClass: 'icon-check-1',
          bg: 'bg-positive',
        }
    }
  }

  if (!status) {
    return null
  }

  return (
    <div
      className={`py-[3.5px] px-[5.5px] text-btn-xs rounded-[4px] inline-block ${getColor()?.text} ${getColor()?.bg}`}
    >
      <span className={`${getColor().iconClass} text-[15px] ${getColor().icon}`} /> {status}
    </div>
  )
}

export default Tag
