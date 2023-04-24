package notifier

type Notifier interface {
	NotifyOrderStatusChange(uri string, obj *OrderStatusChangeReqSt) error
}
