// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/riku179/regisys-server/design
// --out=$(GOPATH)/src/github.com/riku179/regisys-server
// --version=v1.1.0-dirty
//
// API "regisys": CLI Commands
//
// The content of this file is auto-generated, DO NOT MODIFY

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/riku179/regisys-server/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// AddItemsCommand is the command line data structure for the add action of items
	AddItemsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteItemsCommand is the command line data structure for the delete action of items
	DeleteItemsCommand struct {
		// Unique item ID
		ID          int
		PrettyPrint bool
	}

	// ModifyItemsCommand is the command line data structure for the modify action of items
	ModifyItemsCommand struct {
		Payload     string
		ContentType string
		// Unique item ID
		ID          int
		PrettyPrint bool
	}

	// ShowItemsCommand is the command line data structure for the show action of items
	ShowItemsCommand struct {
		// User ID
		User        int
		PrettyPrint bool
	}

	// SigninJWTCommand is the command line data structure for the signin action of jwt
	SigninJWTCommand struct {
		// Is member of MMA
		IsMember string
		// Basic Auth Header
		Authorization string
		PrettyPrint   bool
	}

	// AddOrdersCommand is the command line data structure for the add action of orders
	AddOrdersCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteOrdersCommand is the command line data structure for the delete action of orders
	DeleteOrdersCommand struct {
		// Order ID
		ID          int
		PrettyPrint bool
	}

	// ShowOrdersCommand is the command line data structure for the show action of orders
	ShowOrdersCommand struct {
		// Order date
		Date string
		// Unique user ID
		User        string
		PrettyPrint bool
	}

	// AddUserCommand is the command line data structure for the add action of user
	AddUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ModifyUserCommand is the command line data structure for the modify action of user
	ModifyUserCommand struct {
		Payload     string
		ContentType string
		// Unique user ID
		ID          int
		PrettyPrint bool
	}

	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		// Unique user ID
		ID          int
		PrettyPrint bool
	}

	// ShowListUserCommand is the command line data structure for the showList action of user
	ShowListUserCommand struct {
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add",
		Short: `add action`,
	}
	tmp1 := new(AddItemsCommand)
	sub = &cobra.Command{
		Use:   `items ["/item"]`,
		Short: `Provide items`,
		Long: `Provide items

Payload example:

{
   "item_name": "Ut rerum nobis.",
   "member_price": 2514159059026680118,
   "price": 4715551413911070033,
   "quantity": 24941019248932122
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(AddOrdersCommand)
	sub = &cobra.Command{
		Use:   `orders ["/orders"]`,
		Short: `Provide orders`,
		Long: `Provide orders

Payload example:

{
   "item_id": 520721309930849131,
   "price": 6311483591368200811,
   "quantity": 7849536275300465135,
   "user_id": 5152369185030202978
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(AddUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user"]`,
		Short: ``,
		Long: `

Payload example:

{
   "name": "Consequuntur voluptas aut illo animi.",
   "password": "Aperiam earum."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp4 := new(DeleteItemsCommand)
	sub = &cobra.Command{
		Use:   `items ["/item/ID"]`,
		Short: `Provide items`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(DeleteOrdersCommand)
	sub = &cobra.Command{
		Use:   `orders ["/orders/ID"]`,
		Short: `Provide orders`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "modify",
		Short: `modify action`,
	}
	tmp6 := new(ModifyItemsCommand)
	sub = &cobra.Command{
		Use:   `items ["/item/ID"]`,
		Short: `Provide items`,
		Long: `Provide items

Payload example:

{
   "id": 1299470091090727982,
   "item_name": "Atque et optio molestiae magni quae.",
   "member_price": 6144811010436733541,
   "price": 1718703717495735613,
   "quantity": 5203681815946760459
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(ModifyUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "group": "normal"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp8 := new(ShowItemsCommand)
	sub = &cobra.Command{
		Use:   `items ["/item"]`,
		Short: `Provide items`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(ShowOrdersCommand)
	sub = &cobra.Command{
		Use:   `orders ["/orders"]`,
		Short: `Provide orders`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show-list",
		Short: `Show users list`,
	}
	tmp11 := new(ShowListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/list"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `Create a valid JWT`,
	}
	tmp12 := new(SigninJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/token"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/") {
		fnd = c.Download
		rpath = rpath[1:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the AddItemsCommand command.
func (cmd *AddItemsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/item"
	}
	var payload client.AddGoodsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddItems(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddItemsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteItemsCommand command.
func (cmd *DeleteItemsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/item/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteItems(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteItemsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `Unique item ID`)
}

// Run makes the HTTP request corresponding to the ModifyItemsCommand command.
func (cmd *ModifyItemsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/item/%v", cmd.ID)
	}
	var payload client.ModifyGoodsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ModifyItems(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ModifyItemsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `Unique item ID`)
}

// Run makes the HTTP request corresponding to the ShowItemsCommand command.
func (cmd *ShowItemsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/item"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowItems(ctx, path, intFlagVal("user", cmd.User))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowItemsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var user int
	cc.Flags().IntVar(&cmd.User, "user", user, `User ID`)
}

// Run makes the HTTP request corresponding to the SigninJWTCommand command.
func (cmd *SigninJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/token"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp13 *bool
	if cmd.IsMember != "" {
		var err error
		tmp13, err = boolVal(cmd.IsMember)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--is_member", "err", err)
			return err
		}
	}
	if tmp13 == nil {
		goa.LogError(ctx, "required flag is missing", "flag", "--is_member")
		return fmt.Errorf("required flag is_member is missing")
	}
	resp, err := c.SigninJWT(ctx, path, *tmp13, cmd.Authorization)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var isMember string
	cc.Flags().StringVar(&cmd.IsMember, "is_member", isMember, `Is member of MMA`)
	cc.Flags().StringVar(&cmd.Authorization, "Authorization", "", `Basic Auth Header`)
}

// Run makes the HTTP request corresponding to the AddOrdersCommand command.
func (cmd *AddOrdersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/orders"
	}
	var payload client.AddOrderPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddOrders(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOrdersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteOrdersCommand command.
func (cmd *DeleteOrdersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/orders/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteOrders(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteOrdersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `Order ID`)
}

// Run makes the HTTP request corresponding to the ShowOrdersCommand command.
func (cmd *ShowOrdersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/orders"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp14 *time.Time
	if cmd.Date != "" {
		var err error
		tmp14, err = timeVal(cmd.Date)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *time.Time value", "flag", "--date", "err", err)
			return err
		}
	}
	resp, err := c.ShowOrders(ctx, path, tmp14, stringFlagVal("user", cmd.User))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowOrdersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var date string
	cc.Flags().StringVar(&cmd.Date, "date", date, `Order date`)
	var user string
	cc.Flags().StringVar(&cmd.User, "user", user, `Unique user ID`)
}

// Run makes the HTTP request corresponding to the AddUserCommand command.
func (cmd *AddUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user"
	}
	var payload client.AddUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ModifyUserCommand command.
func (cmd *ModifyUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/user/%v", cmd.ID)
	}
	var payload client.ModifyUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ModifyUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ModifyUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `Unique user ID`)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/user/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id int
	cc.Flags().IntVar(&cmd.ID, "id", id, `Unique user ID`)
}

// Run makes the HTTP request corresponding to the ShowListUserCommand command.
func (cmd *ShowListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/list"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowListUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
