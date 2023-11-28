
export const backendApiURL = {
    public: {
        auth: {
            registration: "public/v1/auth/registration",
            login: "public/v1/auth/login",
            refresh: "public/v1/auth/refresh",
            emailValidation: {
                request: 'public/v1/auth/email-validation/request',
                validate: 'public/v1/auth/email-validation/validate'
            },
            resetPassword: {
                request: 'public/v1/auth/reset-password/request',
                update: 'public/v1/auth/reset-password/update'
            },
            oauth: {
                google: '/public/v1/auth/oauth/google'
            }
        },
        material: {
            get: '/public/v1/material'
        }
    },
    basic: {
        auth: {
            me: "basic/v1/auth/me",
            updatePassword: 'basic/v1/auth/update-password',
            updateAccount: 'basic/v1/auth/update-account'
        },
        question: {
            list: '/basic/v1/question',
            answer: '/basic/v1/question/answer',
            submitAnswer: '/basic/v1/question/submit-answer',
            addRemoveMark: '/basic/v1/question/add-remove-mark',
            getSolution: '/basic/v1/question/solution'
        },
        analytic: {
            attempt: '/basic/v1/analytic/attempt',
            point: '/basic/v1/analytic/point',
            pointList: '/basic/v1/analytic/point-list'
        },
        questionPack: {
            base: 'basic/v1/question-pack'
        },
        product: {
            base: 'basic/v1/product'
        },
    },
    admin: {
        user: {
            list: 'admin/v1/user',
            detail: 'admin/v1/user/'
        },
        role: {
            list: 'admin/v1/role',
            create: 'admin/v1/role',
            detail: 'admin/v1/role/',
            update: 'admin/v1/role/',
            delete: 'admin/v1/role/',
            assign: 'admin/v1/role/assign',
            revoke: 'admin/v1/role/revoke'
        },
        material: {
            list: 'admin/v1/material'
        },
        question: {
            list : 'admin/v1/question',
            option : 'admin/v1/question/option',
            tags: 'admin/v1/question-tag',
            addTags: 'admin/v1/question/add-tags',
            removeTag: 'admin/v1/question/remove-tag',
            solution: 'admin/v1/question/solution',
            questionImage : 'admin/v1/question/question-image',
        },
        analytic: {
            creator: 'admin/v1/analytic'
        },
        questionSolution : {
            base: 'admin/v1/question-solution'
        },
        questionPack: {
            base: 'admin/v1/question-pack'
        }
    },
    contributor: {
        question: {
            base: 'contributor/v1/question',
            tags: {
                base: 'contributor/v1/question/tags',
                remove: 'contributor/v1/question/tags/remove'
            },
            option: {
                base: 'contributor/v1/question/option'
            }
        },
        questionTag: {
            base: 'contributor/v1/question-tag'
        },
        material: {
            base: 'contributor/v1/material'
        }
    }
}
