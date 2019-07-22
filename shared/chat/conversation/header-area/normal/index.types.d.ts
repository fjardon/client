export type Props = {
  badgeNumber?: number
  channelName?: string
  contactNames: {[participant: string]: string}
  muted: boolean
  onBack: () => void
  onOpenFolder?: () => void
  onShowProfile: (user: string) => void
  onToggleInfoPanel: () => void
  onToggleThreadSearch: () => void
  infoPanelOpen: boolean
  teamName?: string
  participants: Array<string>
  participantToDisplayName: {[key: string]: string}
  pendingWaiting: boolean
  smallTeam: boolean
  unMuteConversation: () => void
  username: string
}
