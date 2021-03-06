package enums

type Scope string

// https://www.admitad.com/en/developers/doc/api_ru/auth/auth-rights/
const (
	PublicData               Scope = "public_data"
	Websites                 Scope = "websites"
	ManageWebsites           Scope = "manage_websites"
	AdvCampaigns             Scope = "advcampaigns"
	AdvCampaignsForSite      Scope = "advcampaigns_for_website"
	ManageAdvCampaigns       Scope = "manage_advcampaigns"
	Banners                  Scope = "banners"
	Landings                 Scope = "landings"
	BannersForWebsite        Scope = "banners_for_website"
	Payments                 Scope = "payments"
	ManagePayments           Scope = "manage_payments"
	Announcements            Scope = "announcements"
	Referrals                Scope = "referrals"
	Coupons                  Scope = "coupons"
	CouponsForSIte           Scope = "coupons_for_website"
	PrivateData              Scope = "private_data"
	Tickets                  Scope = "tickets"
	ManageTickets            Scope = "manage_tickets"
	PrivateDataEmail         Scope = "private_data_email"
	PrivateDataPhone         Scope = "private_data_phone"
	PrivateDataBalance       Scope = "private_data_balance"
	ValidateLinks            Scope = "validate_links"
	DeepLinkGenerator        Scope = "deeplink_generator"
	Statistics               Scope = "statistics"
	OptCodes                 Scope = "opt_codes"
	ManageOptCodes           Scope = "manage_opt_codes"
	WebmasterReTag           Scope = "webmaster_retag"
	ManageWebmastersReTag    Scope = "manage_webmaster_retag"
	BrokenLinks              Scope = "broken_links"
	ManageBrokenLinks        Scope = "manage_broken_links"
	LostOrders               Scope = "lost_orders"
	ManageLostOrders         Scope = "manage_lost_orders"
	BrokerApplications       Scope = "broker_application"
	ManageBrokerApplications Scope = "manage_broker_application"
	OfflineSales             Scope = "offline_sales"
	OfflineReceipts          Scope = "offline_receipts"
	ManageOfflineReceipts    Scope = "manage_offline_receipts"
)
