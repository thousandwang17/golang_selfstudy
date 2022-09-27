package main

type meditorer interface {
	canArrived(t tainer) bool
	notifyAboutDepature()
}
