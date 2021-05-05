import React, { cloneElement } from 'react'
import {
  Button,
  sanitizeListRestProps,
  TopToolbar,
  useTranslate,
} from 'react-admin'
import {
  ButtonGroup,
  useMediaQuery,
  Typography,
  makeStyles,
} from '@material-ui/core'
import ViewHeadlineIcon from '@material-ui/icons/ViewHeadline'
import ViewModuleIcon from '@material-ui/icons/ViewModule'
import { useDispatch, useSelector } from 'react-redux'
import { albumViewGrid, albumViewList } from '../actions'
import ToggleFieldsMenu from '../common/ToggleFieldsMenu'

const useStyles = makeStyles({
  title: { margin: '1rem' },
  buttonGroup: { width: '100%', justifyContent: 'center' },
  leftButton: { paddingRight: '0.5rem' },
  rightButton: { paddingLeft: '0.5rem' },
})

const AlbumViewToggler = React.forwardRef(({}, ref) => {
  const dispatch = useDispatch()
  const albumView = useSelector((state) => state.albumView)
  const classes = useStyles()
  const translate = useTranslate()
  return (
    <div ref={ref}>
      <Typography className={classes.title}>
        {translate('ra.toggleFieldsMenu.layout')}
      </Typography>
      <ButtonGroup
        variant="text"
        color="primary"
        aria-label="text primary button group"
        className={classes.buttonGroup}
      >
        <Button
          size="small"
          className={classes.leftButton}
          label="Grid"
          color={albumView.grid ? 'primary' : 'secondary'}
          onClick={() => dispatch(albumViewGrid())}
        >
          <ViewModuleIcon fontSize="inherit" />
        </Button>
        <Button
          size="small"
          className={classes.rightButton}
          label="Table"
          color={albumView.grid ? 'secondary' : 'primary'}
          onClick={() => dispatch(albumViewList())}
        >
          <ViewHeadlineIcon fontSize="inherit" />
        </Button>
      </ButtonGroup>
    </div>
  )
})

const AlbumListActions = ({
  currentSort,
  className,
  resource,
  filters,
  displayedFilters,
  filterValues,
  permanentFilter,
  exporter,
  basePath,
  selectedIds,
  onUnselectItems,
  showFilter,
  maxResults,
  total,
  fullWidth,
  ...rest
}) => {
  const isSmall = useMediaQuery((theme) => theme.breakpoints.up('sm'))
  return (
    <TopToolbar className={className} {...sanitizeListRestProps(rest)}>
      {filters &&
        cloneElement(filters, {
          resource,
          showFilter,
          displayedFilters,
          filterValues,
          context: 'button',
        })}
      {isSmall && (
        <ToggleFieldsMenu resource="album" TopBarComponent={AlbumViewToggler} />
      )}
    </TopToolbar>
  )
}

AlbumListActions.defaultProps = {
  selectedIds: [],
  onUnselectItems: () => null,
}

export default AlbumListActions
