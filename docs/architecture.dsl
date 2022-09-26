workspace {
    model {
        basicUser = person "Basic User"
        accountingUser = person "Accounting User"
        adminUser = person "Admin User"
        
        authService = softwareSystem "AuthService" {
            authServiceAPIApplication = container "API service" {
                authRolesManagementComponent = component "Role management" "Add / Delete user roles"
                authUserManagementComponent = component "User management" "Register user / Assign user roles / List users"
                authTokenGenerationComponent = component "Token generation" "Generate JWT with user roles"
            }
            
            authServiceDB = container "Database" "Stores information about users, roles and auth tokens" "" "Database"
        }
        
        tasksService = softwareSystem "TasksService" {
            tasksServiceAPIApplication = container "API service" {
                tasksAPITasksManagementComponent = component "Tasks management" "Create / Update / List tasks"
                tasksAPITasksDistributionComponent = component "Tasks distribution" "Reassign all tasks"
            }
            
            tasksServiceDB = container "Database" "Stores information about tasks" "" "Database"
            tasksServiceWebApplication = container "Web application" "UI"
        }
        
        accountingService = softwareSystem "AccountingService" {
            accountingServiceAPIApplication = container "API service" {
                acountingAPITransactionManagementComponent = component "Transaction management" "Widraw / deposit on user account, audit log"
                acountingAPIPayoffManagementComponent = component "Payoff management" "Pay off money to the users"
                accountingAPIReportingComponent = component "Reporting" "Dashboard for current state of the account(s), historical changes"
            }
            accountingServiceDB = container "Database" "Stores information" "" "Database"
            accountingServiceWebApplication = container "Web application" "UI"
        }
        
        
        analyticService = softwareSystem "AnaliticService" 
        
        # person to service
        basicUser -> tasksService "Uses"
        basicUser -> accountingService "Uses"
        
        accountingUser -> tasksService "Uses"
        accountingUser -> accountingService "Uses"
        
        adminUser -> tasksService "Uses"
        adminUser -> accountingService "Uses"
        adminUser -> analyticService "Uses"
        
        # service to service
        analyticService -> authService "Get user role"
        tasksService -> authService "Get user role"
        accountingService -> authService "Get user role"
        
        # containers
        tasksServiceWebApplication -> authServiceAPIApplication "Get JWT" "JSON/HTTPS"
        tasksServiceAPIApplication -> tasksServiceDB "Reads from and writes to"
        tasksServiceWebApplication -> tasksServiceAPIApplication "Makes API calls to" "JSON/HTTPS"
        tasksServiceAPIApplication -> authServiceAPIApplication "Get public key to validate JWT" "JSON/HTTPS"
    }

    views {
        systemLandscape {
            include *
            autolayout lr
        }
        
        systemcontext authService "AuthSystemContext" {
            include *
            
            autoLayout
        }
        
        container authService "AuthContainers" {
            include *

            autoLayout
        }
        
        component authServiceAPIApplication "AuthAPIComponents" {
            include *
            autoLayout
        }
        
        systemcontext tasksService "TasksSystemContext" {
            include *
            
            autoLayout
        }
        
        container tasksService "TasksContainers" {
            include *

            autoLayout
        }
        
        component tasksServiceAPIApplication "TasksAPIComponents" {
            include *
            autoLayout
        }
        
        systemcontext accountingService "AcountingSystemContext" {
            include *
            
            autoLayout
        }
        
        container accountingService "AcountingContainers" {
            include *

            autoLayout
        }
        
        component accountingServiceAPIApplication "AcountingAPIComponents" {
            include *
            autoLayout
        }
        
        styles {
            element "Database" {
                shape Cylinder
            }
        }
        theme default
    }
}