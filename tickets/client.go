package tickets

import (
	client "github.com/lueurxax/go-admitad-client"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
	"strconv"
)

type Ticket struct {
	baseClient *client.AClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/tickets/tickets-list/
func (c *Ticket) TicketByID(id int) error {
	answer := new(responses.Ticket)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
		SetResult(answer).
		SetError(errResponse).
		Post("/tickets/" + strconv.Itoa(id))

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/tickets/tickets-list/
func (c *Ticket) Ticket(ticket requests.Ticket) error {
	answer := new(responses.Ticket)
	errResponse := new(responses.ErrorResponse)

	params := map[string]string{}

	if ticket.Status != 0 {
		params["status"] = strconv.Itoa(ticket.Status)
	}

	if ticket.Limit != 0 {
		params["limit"] = strconv.Itoa(ticket.Limit)
	}

	if ticket.Offset != 0 {
		params["offset"] = strconv.Itoa(ticket.Offset)
	}

	if ticket.DateEnd != nil {
		params["date_start"] = ticket.DateStart.Format("02.01.2006")
	}

	if ticket.DateEnd != nil {
		params["date_end"] = ticket.DateEnd.Format("02.01.2006")
	}

	resp, err := c.baseClient.R().EnableTrace().
		SetQueryParams(params).
		SetAuthToken(c.baseClient.Token).
		SetResult(answer).
		SetError(errResponse).
		Post("/tickets/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/tickets/tickets-create/
// Все значения в структуре requests.CreateTicket обязательные
func (c *Ticket) CreateTicket(ticket requests.CreateTicket) error {
	answer := new(responses.Ticket)
	errResponse := new(responses.ErrorResponse)
	params := map[string]string{}

	if ticket.AdvCampaign != 0 {
		params["advcampaign"] = strconv.Itoa(ticket.AdvCampaign)
	}

	if ticket.Category != 0 {
		params["category"] = strconv.Itoa(ticket.Category)
	}

	if ticket.Priority != 0 {
		params["priority"] = strconv.Itoa(ticket.Priority)
	}

	if ticket.Text != "" {
		params["text"] = ticket.Text
	}

	if ticket.Subject != "" {
		params["subject"] = ticket.Subject
	}

	resp, err := c.baseClient.R().EnableTrace().
		SetQueryParams(params).
		SetAuthToken(c.baseClient.Token).
		SetResult(answer).
		SetError(errResponse).
		Post("/tickets/create/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}
